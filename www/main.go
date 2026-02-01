package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sort"
	"syscall"
	"time"
)

func main() {
	resolversDir := flag.String("resolvers", "../v3", "Path to the resolvers directory")
	dbPath := flag.String("db", "resolvers.db", "Path to SQLite database")
	webAddr := flag.String("addr", ":8080", "Web server address")
	testInterval := flag.Duration("interval", 5*time.Minute, "Test interval")
	concurrency := flag.Int("concurrency", 50, "Number of concurrent tests")
	timeout := flag.Duration("timeout", 10*time.Second, "Timeout for each test")
	runOnce := flag.Bool("once", false, "Run tests once and exit")
	rebuildStats := flag.Bool("rebuild-stats", false, "Rebuild stats from test_results and exit")
	removeResolver := flag.String("remove", "", "Remove a resolver by name and exit")
	clearErrors := flag.String("clear-errors", "", "Clear all errors for a resolver (mark as successful) and exit")
	hideLatency := flag.Bool("hide-latency", false, "Hide latency column in web interface")
	minReliability := flag.Float64("min-reliability", -1, "Remove resolvers with reliability below this percentage and exit")
	showFailures := flag.Bool("failures", false, "Print failing resolvers/relays with their error messages and exit")
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	db, err := NewDB(*dbPath)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// Handle explicit rebuild request
	if *rebuildStats {
		log.Println("Rebuilding resolver stats...")
		if err := db.RebuildStats(); err != nil {
			log.Fatalf("Failed to rebuild stats: %v", err)
		}
		log.Println("Stats rebuilt successfully")
		return
	}

	// Handle resolver removal
	if *removeResolver != "" {
		log.Printf("Removing resolver %q...", *removeResolver)
		if err := db.RemoveResolver(*removeResolver); err != nil {
			log.Fatalf("Failed to remove resolver: %v", err)
		}
		log.Printf("Resolver %q removed successfully", *removeResolver)
		return
	}

	// Handle clearing errors for a resolver
	if *clearErrors != "" {
		log.Printf("Clearing errors for resolver %q...", *clearErrors)
		updated, err := db.ClearResolverErrors(*clearErrors)
		if err != nil {
			log.Fatalf("Failed to clear errors: %v", err)
		}
		log.Printf("Cleared %d failed tests for resolver %q", updated, *clearErrors)
		return
	}

	// Handle removing unreliable resolvers
	if *minReliability >= 0 {
		log.Printf("Removing resolvers with reliability below %.2f%%...", *minReliability)
		removed, err := db.RemoveUnreliableResolvers(*minReliability)
		if err != nil {
			log.Fatalf("Failed to remove unreliable resolvers: %v", err)
		}
		if len(removed) == 0 {
			log.Println("No resolvers removed")
		} else {
			log.Printf("Removed %d resolvers: %v", len(removed), removed)
		}
		return
	}

	// Handle showing failures
	if *showFailures {
		stats, err := db.GetAllStats()
		if err != nil {
			log.Fatalf("Failed to get stats: %v", err)
		}
		failures := filterFailures(stats)
		if len(failures) == 0 {
			log.Println("No failing resolvers/relays")
			return
		}
		for _, s := range failures {
			fmt.Printf("%s (%s): %s\n", s.Name, s.Type, s.LastError)
		}
		return
	}

	// Auto-rebuild if resolver_stats is empty but test_results has data
	if err := db.RebuildStatsIfNeeded(); err != nil {
		log.Printf("Warning: failed to check/rebuild stats: %v", err)
	}

	parser := NewParser(*resolversDir)
	tester := NewTester(db, *concurrency, *timeout)

	runTests := func() {
		startTime := time.Now()
		log.Printf("Starting test run at %s", startTime.Format(time.RFC3339))

		resolvers, err := parser.ParseAll()
		if err != nil {
			log.Printf("Failed to parse resolvers: %v", err)
			return
		}
		log.Printf("Parsed %d resolvers/relays", len(resolvers))
		tester.TestAll(resolvers)

		removed, err := db.RemoveStaleResolvers(7 * 24 * time.Hour)
		if err != nil {
			log.Printf("Failed to remove stale resolvers: %v", err)
		} else if len(removed) > 0 {
			log.Printf("Removed %d stale resolvers: %v", len(removed), removed)
		}

		// Prune test results older than 1 week
		pruned, err := db.PruneOldTests(7 * 24 * time.Hour)
		if err != nil {
			log.Printf("Failed to prune old tests: %v", err)
		} else if pruned > 0 {
			log.Printf("Pruned %d old test results", pruned)
		}

		// Log database state after test run
		count, lastTest, err := db.GetTestCount()
		if err != nil {
			log.Printf("Failed to get test count: %v", err)
		} else {
			log.Printf("Database: %d total tests, last at %s", count, lastTest.Format(time.RFC3339))
		}

		log.Printf("Test run completed in %v", time.Since(startTime))
	}

	if *runOnce {
		runTests()
		return
	}

	web := NewWebServer(db, *webAddr, *hideLatency)
	go web.Start()

	ticker := time.NewTicker(*testInterval)
	defer ticker.Stop()

	runTests()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-ticker.C:
			runTests()
		case <-sigChan:
			log.Println("Shutting down...")
			return
		}
	}
}

// filterFailures returns resolvers where the last test failed.
// A resolver is considered failing if it has a LastFail time that is
// more recent than LastSuccess (or if it has never succeeded).
func filterFailures(stats []ResolverStats) []ResolverStats {
	var failures []ResolverStats
	for _, s := range stats {
		if s.LastError == "" {
			continue
		}
		// Consider it failing if last_fail > last_success
		if s.LastFail != nil {
			if s.LastSuccess == nil || s.LastFail.After(*s.LastSuccess) {
				failures = append(failures, s)
			}
		}
	}
	sort.Slice(failures, func(i, j int) bool {
		if failures[i].Type != failures[j].Type {
			return failures[i].Type < failures[j].Type
		}
		return failures[i].Name < failures[j].Name
	})
	return failures
}

package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
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
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	db, err := NewDB(*dbPath)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	parser := NewParser(*resolversDir)
	tester := NewTester(db, *concurrency, *timeout)

	if *runOnce {
		resolvers, err := parser.ParseAll()
		if err != nil {
			log.Fatalf("Failed to parse resolvers: %v", err)
		}
		log.Printf("Parsed %d resolvers/relays", len(resolvers))
		tester.TestAll(resolvers)
		return
	}

	web := NewWebServer(db, *webAddr)
	go web.Start()

	ticker := time.NewTicker(*testInterval)
	defer ticker.Stop()

	runTests := func() {
		resolvers, err := parser.ParseAll()
		if err != nil {
			log.Printf("Failed to parse resolvers: %v", err)
			return
		}
		log.Printf("Parsed %d resolvers/relays", len(resolvers))
		tester.TestAll(resolvers)
	}

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

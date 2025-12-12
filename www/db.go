package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	db *sql.DB
}

type TestResult struct {
	ID          int64
	Name        string
	Type        string
	Stamp       string
	Success     bool
	RTT         int64
	Error       string
	TestedAt    time.Time
	Description string
	SourceFile  string
}

type ResolverStats struct {
	Name           string
	Type           string
	Description    string
	SourceFile     string
	Stamp          string
	TotalTests     int
	SuccessCount   int
	FailCount      int
	AvgRTT         float64
	MinRTT         int64
	MaxRTT         int64
	LastSuccess    *time.Time
	LastFail       *time.Time
	LastTestedAt   time.Time
	LastError      string
	ReliabilityPct float64
}

func NewDB(path string) (*DB, error) {
	db, err := sql.Open("sqlite3", path+"?_journal_mode=WAL&_busy_timeout=5000&_synchronous=NORMAL")
	if err != nil {
		return nil, err
	}

	// Verify connection works
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	if err := createTables(db); err != nil {
		db.Close()
		return nil, err
	}

	return &DB{db: db}, nil
}

func createTables(db *sql.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS resolvers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		type TEXT NOT NULL,
		description TEXT,
		source_file TEXT,
		UNIQUE(name, type)
	);

	CREATE TABLE IF NOT EXISTS test_results (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		resolver_id INTEGER NOT NULL,
		stamp TEXT NOT NULL,
		success INTEGER NOT NULL,
		rtt_ms INTEGER,
		error TEXT,
		tested_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (resolver_id) REFERENCES resolvers(id)
	);

	-- For GetTestCount: MAX(tested_at)
	CREATE INDEX IF NOT EXISTS idx_test_results_tested_at ON test_results(tested_at);

	-- For RebuildStats: fetching latest stamp per resolver
	CREATE INDEX IF NOT EXISTS idx_test_results_resolver_tested_at ON test_results(resolver_id, tested_at DESC);

	-- For RebuildStats (last error) and RemoveStaleResolvers
	CREATE INDEX IF NOT EXISTS idx_test_results_resolver_success ON test_results(resolver_id, success, tested_at DESC);

	-- Pre-computed stats table for fast queries
	CREATE TABLE IF NOT EXISTS resolver_stats (
		resolver_id INTEGER PRIMARY KEY,
		stamp TEXT,
		total_tests INTEGER DEFAULT 0,
		success_count INTEGER DEFAULT 0,
		fail_count INTEGER DEFAULT 0,
		avg_rtt REAL DEFAULT 0,
		min_rtt INTEGER DEFAULT 0,
		max_rtt INTEGER DEFAULT 0,
		rtt_sum INTEGER DEFAULT 0,
		last_success DATETIME,
		last_fail DATETIME,
		last_tested DATETIME,
		last_error TEXT,
		reliability_pct REAL DEFAULT 0,
		FOREIGN KEY (resolver_id) REFERENCES resolvers(id)
	);
	`
	_, err := db.Exec(schema)
	return err
}

func (d *DB) Close() error {
	return d.db.Close()
}

func (d *DB) UpsertResolver(name, typ, description, sourceFile string) (int64, error) {
	_, err := d.db.Exec(`
		INSERT INTO resolvers (name, type, description, source_file)
		VALUES (?, ?, ?, ?)
		ON CONFLICT(name, type) DO UPDATE SET
			description = excluded.description,
			source_file = excluded.source_file
	`, name, typ, description, sourceFile)
	if err != nil {
		return 0, err
	}

	// Always query for the ID after upsert. LastInsertId() is unreliable with
	// ON CONFLICT DO UPDATE - it returns the last insert rowid from any prior
	// insert in the session, not the ID of the updated row.
	var id int64
	row := d.db.QueryRow("SELECT id FROM resolvers WHERE name = ? AND type = ?", name, typ)
	err = row.Scan(&id)
	return id, err
}

func (d *DB) RecordTest(resolverID int64, stamp string, success bool, rttMs int64, errMsg string) error {
	result, err := d.db.Exec(`
		INSERT INTO test_results (resolver_id, stamp, success, rtt_ms, error)
		VALUES (?, ?, ?, ?, ?)
	`, resolverID, stamp, success, rttMs, errMsg)
	if err != nil {
		return err
	}
	// Verify the insert actually happened
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("no rows inserted for resolver %d", resolverID)
	}

	// Update pre-computed stats
	return d.updateResolverStats(resolverID, stamp, success, rttMs, errMsg)
}

func (d *DB) updateResolverStats(resolverID int64, stamp string, success bool, rttMs int64, errMsg string) error {
	now := time.Now().Format("2006-01-02 15:04:05")

	// Upsert into resolver_stats with incremental updates
	var query string
	var args []interface{}

	if success {
		query = `
			INSERT INTO resolver_stats (resolver_id, stamp, total_tests, success_count, fail_count,
				rtt_sum, avg_rtt, min_rtt, max_rtt, last_success, last_tested, reliability_pct)
			VALUES (?, ?, 1, 1, 0, ?, ?, ?, ?, ?, ?, 100.0)
			ON CONFLICT(resolver_id) DO UPDATE SET
				stamp = excluded.stamp,
				total_tests = total_tests + 1,
				success_count = success_count + 1,
				rtt_sum = rtt_sum + ?,
				avg_rtt = CAST(rtt_sum + ? AS REAL) / (success_count + 1),
				min_rtt = CASE WHEN min_rtt = 0 OR ? < min_rtt THEN ? ELSE min_rtt END,
				max_rtt = CASE WHEN ? > max_rtt THEN ? ELSE max_rtt END,
				last_success = ?,
				last_tested = ?,
				reliability_pct = CAST(success_count + 1 AS REAL) / (total_tests + 1) * 100.0
		`
		args = []interface{}{
			resolverID, stamp, rttMs, rttMs, rttMs, rttMs, now, now, // INSERT values
			rttMs, rttMs, rttMs, rttMs, rttMs, rttMs, now, now, // UPDATE values
		}
	} else {
		query = `
			INSERT INTO resolver_stats (resolver_id, stamp, total_tests, success_count, fail_count,
				last_fail, last_tested, last_error, reliability_pct)
			VALUES (?, ?, 1, 0, 1, ?, ?, ?, 0.0)
			ON CONFLICT(resolver_id) DO UPDATE SET
				stamp = excluded.stamp,
				total_tests = total_tests + 1,
				fail_count = fail_count + 1,
				last_fail = ?,
				last_tested = ?,
				last_error = ?,
				reliability_pct = CAST(success_count AS REAL) / (total_tests + 1) * 100.0
		`
		args = []interface{}{
			resolverID, stamp, now, now, errMsg, // INSERT values
			now, now, errMsg, // UPDATE values
		}
	}

	_, err := d.db.Exec(query, args...)
	return err
}

func (d *DB) GetAllStats() ([]ResolverStats, error) {
	// Use pre-computed stats table for fast queries
	// No ORDER BY here - web.go sorts in Go based on user preference
	rows, err := d.db.Query(`
		SELECT
			r.name,
			r.type,
			r.description,
			r.source_file,
			s.stamp,
			s.total_tests,
			s.success_count,
			s.fail_count,
			s.avg_rtt,
			s.min_rtt,
			s.max_rtt,
			s.last_success,
			s.last_fail,
			s.last_tested,
			s.last_error,
			s.reliability_pct
		FROM resolvers r
		LEFT JOIN resolver_stats s ON r.id = s.resolver_id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []ResolverStats
	for rows.Next() {
		var s ResolverStats
		var lastSuccess, lastFail, lastTested sql.NullString
		var lastError sql.NullString
		var description, sourceFile, stamp sql.NullString
		var totalTests, successCount, failCount sql.NullInt64
		var avgRTT, reliabilityPct sql.NullFloat64
		var minRTT, maxRTT sql.NullInt64

		err := rows.Scan(
			&s.Name, &s.Type, &description, &sourceFile, &stamp,
			&totalTests, &successCount, &failCount,
			&avgRTT, &minRTT, &maxRTT,
			&lastSuccess, &lastFail, &lastTested, &lastError, &reliabilityPct,
		)
		if err != nil {
			return nil, err
		}

		s.Description = description.String
		s.SourceFile = sourceFile.String
		s.Stamp = stamp.String
		s.LastError = lastError.String
		s.TotalTests = int(totalTests.Int64)
		s.SuccessCount = int(successCount.Int64)
		s.FailCount = int(failCount.Int64)
		s.AvgRTT = avgRTT.Float64
		s.MinRTT = minRTT.Int64
		s.MaxRTT = maxRTT.Int64
		s.ReliabilityPct = reliabilityPct.Float64

		if lastSuccess.Valid {
			t, _ := time.Parse("2006-01-02 15:04:05", lastSuccess.String)
			s.LastSuccess = &t
		}
		if lastFail.Valid {
			t, _ := time.Parse("2006-01-02 15:04:05", lastFail.String)
			s.LastFail = &t
		}
		if lastTested.Valid {
			s.LastTestedAt, _ = time.Parse("2006-01-02 15:04:05", lastTested.String)
		}

		stats = append(stats, s)
	}

	return stats, rows.Err()
}

// RebuildStatsIfNeeded rebuilds stats only if the table is empty but test_results has data.
func (d *DB) RebuildStatsIfNeeded() error {
	var statsCount, testCount int
	if err := d.db.QueryRow("SELECT COUNT(*) FROM resolver_stats").Scan(&statsCount); err != nil {
		return err
	}
	if err := d.db.QueryRow("SELECT COUNT(*) FROM test_results").Scan(&testCount); err != nil {
		return err
	}

	if statsCount == 0 && testCount > 0 {
		return d.RebuildStats()
	}
	return nil
}

// RebuildStats recomputes all resolver_stats from test_results.
// Call this once to migrate existing data, or to fix any inconsistencies.
func (d *DB) RebuildStats() error {
	_, err := d.db.Exec(`
		DELETE FROM resolver_stats;

		INSERT INTO resolver_stats (
			resolver_id, stamp, total_tests, success_count, fail_count,
			rtt_sum, avg_rtt, min_rtt, max_rtt,
			last_success, last_fail, last_tested, last_error, reliability_pct
		)
		SELECT
			r.id,
			(SELECT stamp FROM test_results t3 WHERE t3.resolver_id = r.id ORDER BY t3.tested_at DESC LIMIT 1),
			COUNT(t.id),
			SUM(CASE WHEN t.success = 1 THEN 1 ELSE 0 END),
			SUM(CASE WHEN t.success = 0 THEN 1 ELSE 0 END),
			COALESCE(SUM(CASE WHEN t.success = 1 THEN t.rtt_ms ELSE 0 END), 0),
			COALESCE(AVG(CASE WHEN t.success = 1 THEN t.rtt_ms END), 0),
			COALESCE(MIN(CASE WHEN t.success = 1 THEN t.rtt_ms END), 0),
			COALESCE(MAX(CASE WHEN t.success = 1 THEN t.rtt_ms END), 0),
			MAX(CASE WHEN t.success = 1 THEN t.tested_at END),
			MAX(CASE WHEN t.success = 0 THEN t.tested_at END),
			MAX(t.tested_at),
			(SELECT error FROM test_results t2 WHERE t2.resolver_id = r.id AND t2.success = 0 ORDER BY t2.tested_at DESC LIMIT 1),
			CASE WHEN COUNT(t.id) > 0 THEN CAST(SUM(CASE WHEN t.success = 1 THEN 1 ELSE 0 END) AS REAL) / COUNT(t.id) * 100.0 ELSE 0 END
		FROM resolvers r
		LEFT JOIN test_results t ON r.id = t.resolver_id
		GROUP BY r.id
	`)
	return err
}

// RemoveStaleResolvers removes resolvers that haven't had a successful response
// in the given duration. Only removes resolvers that have had at least one
// successful test in the past (to avoid removing newly added resolvers).
// Returns the names of removed resolvers.
func (d *DB) RemoveStaleResolvers(noSuccessSince time.Duration) ([]string, error) {
	cutoff := time.Now().Add(-noSuccessSince)

	// Find resolvers with no successful test after cutoff,
	// but that have had at least one successful test before
	rows, err := d.db.Query(`
		SELECT r.id, r.name
		FROM resolvers r
		WHERE NOT EXISTS (
			SELECT 1 FROM test_results t
			WHERE t.resolver_id = r.id
			AND t.success = 1
			AND t.tested_at > ?
		)
		AND EXISTS (
			SELECT 1 FROM test_results t2
			WHERE t2.resolver_id = r.id
			AND t2.success = 1
		)
	`, cutoff)
	if err != nil {
		return nil, err
	}

	var idsToDelete []int64
	var names []string
	for rows.Next() {
		var id int64
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			rows.Close()
			return nil, err
		}
		idsToDelete = append(idsToDelete, id)
		names = append(names, name)
	}
	rows.Close()

	if len(idsToDelete) == 0 {
		return nil, nil
	}

	// Delete test results, stats, and resolvers
	for _, id := range idsToDelete {
		if _, err := d.db.Exec("DELETE FROM test_results WHERE resolver_id = ?", id); err != nil {
			return names, err
		}
		if _, err := d.db.Exec("DELETE FROM resolver_stats WHERE resolver_id = ?", id); err != nil {
			return names, err
		}
		if _, err := d.db.Exec("DELETE FROM resolvers WHERE id = ?", id); err != nil {
			return names, err
		}
	}

	return names, nil
}

// GetTestCount returns the total number of test results and the timestamp of the most recent one
func (d *DB) GetTestCount() (count int64, lastTest time.Time, err error) {
	var lastTestStr sql.NullString
	err = d.db.QueryRow(`
		SELECT COUNT(*), MAX(tested_at) FROM test_results
	`).Scan(&count, &lastTestStr)
	if err != nil {
		return 0, time.Time{}, err
	}
	if lastTestStr.Valid {
		lastTest, _ = time.Parse("2006-01-02 15:04:05", lastTestStr.String)
	}
	return count, lastTest, nil
}

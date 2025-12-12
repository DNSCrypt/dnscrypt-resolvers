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

	CREATE INDEX IF NOT EXISTS idx_test_results_resolver ON test_results(resolver_id);
	CREATE INDEX IF NOT EXISTS idx_test_results_tested_at ON test_results(tested_at);

	-- Composite index for correlated subqueries that fetch latest stamp/error
	CREATE INDEX IF NOT EXISTS idx_test_results_resolver_tested_at ON test_results(resolver_id, tested_at DESC);

	-- Composite index for success-based aggregations and filtering
	CREATE INDEX IF NOT EXISTS idx_test_results_resolver_success ON test_results(resolver_id, success, tested_at DESC);
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
	return nil
}

func (d *DB) GetAllStats() ([]ResolverStats, error) {
	rows, err := d.db.Query(`
		SELECT
			r.name,
			r.type,
			r.description,
			r.source_file,
			(SELECT stamp FROM test_results t3 WHERE t3.resolver_id = r.id ORDER BY t3.tested_at DESC LIMIT 1) as stamp,
			COUNT(t.id) as total_tests,
			SUM(CASE WHEN t.success = 1 THEN 1 ELSE 0 END) as success_count,
			SUM(CASE WHEN t.success = 0 THEN 1 ELSE 0 END) as fail_count,
			COALESCE(AVG(CASE WHEN t.success = 1 THEN t.rtt_ms END), 0) as avg_rtt,
			COALESCE(MIN(CASE WHEN t.success = 1 THEN t.rtt_ms END), 0) as min_rtt,
			COALESCE(MAX(CASE WHEN t.success = 1 THEN t.rtt_ms END), 0) as max_rtt,
			MAX(CASE WHEN t.success = 1 THEN t.tested_at END) as last_success,
			MAX(CASE WHEN t.success = 0 THEN t.tested_at END) as last_fail,
			MAX(t.tested_at) as last_tested,
			(SELECT error FROM test_results t2 WHERE t2.resolver_id = r.id AND t2.success = 0 ORDER BY t2.tested_at DESC LIMIT 1) as last_error
		FROM resolvers r
		LEFT JOIN test_results t ON r.id = t.resolver_id
		GROUP BY r.id
		ORDER BY
			CASE WHEN COUNT(t.id) = 0 THEN 1 ELSE 0 END,
			CAST(SUM(CASE WHEN t.success = 1 THEN 1 ELSE 0 END) AS FLOAT) / NULLIF(COUNT(t.id), 0) DESC,
			AVG(CASE WHEN t.success = 1 THEN t.rtt_ms END) ASC
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

		err := rows.Scan(
			&s.Name, &s.Type, &description, &sourceFile, &stamp,
			&s.TotalTests, &s.SuccessCount, &s.FailCount,
			&s.AvgRTT, &s.MinRTT, &s.MaxRTT,
			&lastSuccess, &lastFail, &lastTested, &lastError,
		)
		if err != nil {
			return nil, err
		}

		s.Description = description.String
		s.SourceFile = sourceFile.String
		s.Stamp = stamp.String
		s.LastError = lastError.String

		if s.TotalTests > 0 {
			s.ReliabilityPct = float64(s.SuccessCount) / float64(s.TotalTests) * 100
		}

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

	// Delete test results and resolvers
	for _, id := range idsToDelete {
		if _, err := d.db.Exec("DELETE FROM test_results WHERE resolver_id = ?", id); err != nil {
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

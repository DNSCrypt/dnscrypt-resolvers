package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"
)

type WebServer struct {
	db   *DB
	addr string
}

func NewWebServer(db *DB, addr string) *WebServer {
	return &WebServer{db: db, addr: addr}
}

func (w *WebServer) Start() {
	http.HandleFunc("/", w.handleIndex)
	http.HandleFunc("/api/stats", w.handleAPIStats)
	http.HandleFunc("/api/stats/", w.handleAPIStatsByType)

	log.Printf("Web server starting on %s", w.addr)
	if err := http.ListenAndServe(w.addr, nil); err != nil {
		log.Fatalf("Web server failed: %v", err)
	}
}

func (w *WebServer) handleIndex(rw http.ResponseWriter, r *http.Request) {
	stats, err := w.db.GetAllStats()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	typeFilter := r.URL.Query().Get("type")
	sortBy := r.URL.Query().Get("sort")
	sortOrder := r.URL.Query().Get("order")
	if sortOrder == "" {
		sortOrder = "asc"
	}
	desc := sortOrder == "desc"

	var filtered []ResolverStats
	if typeFilter != "" {
		for _, s := range stats {
			if s.Type == typeFilter {
				filtered = append(filtered, s)
			}
		}
	} else {
		filtered = stats
	}

	switch sortBy {
	case "name":
		sort.Slice(filtered, func(i, j int) bool {
			if desc {
				return filtered[i].Name > filtered[j].Name
			}
			return filtered[i].Name < filtered[j].Name
		})
	case "type":
		sort.Slice(filtered, func(i, j int) bool {
			if desc {
				return filtered[i].Type > filtered[j].Type
			}
			return filtered[i].Type < filtered[j].Type
		})
	case "rtt":
		sort.Slice(filtered, func(i, j int) bool {
			if desc {
				return filtered[i].AvgRTT > filtered[j].AvgRTT
			}
			return filtered[i].AvgRTT < filtered[j].AvgRTT
		})
	case "reliability":
		sort.Slice(filtered, func(i, j int) bool {
			if desc {
				return filtered[i].ReliabilityPct < filtered[j].ReliabilityPct
			}
			return filtered[i].ReliabilityPct > filtered[j].ReliabilityPct
		})
	}

	types := make(map[string]int)
	for _, s := range stats {
		types[s.Type]++
	}

	data := struct {
		Stats       []ResolverStats
		Types       map[string]int
		TypeFilter  string
		SortBy      string
		SortOrder   string
		LastUpdated time.Time
	}{
		Stats:       filtered,
		Types:       types,
		TypeFilter:  typeFilter,
		SortBy:      sortBy,
		SortOrder:   sortOrder,
		LastUpdated: time.Now(),
	}

	tmpl, err := template.New("index").Funcs(template.FuncMap{
		"formatTime": func(t time.Time) string {
			if t.IsZero() {
				return "Never"
			}
			return t.Format("2006-01-02 15:04:05")
		},
		"formatTimePtr": func(t *time.Time) string {
			if t == nil {
				return "Never"
			}
			return t.Format("2006-01-02 15:04:05")
		},
		"formatDuration": func(t time.Time) string {
			if t.IsZero() {
				return "Never"
			}
			d := time.Since(t)
			if d < time.Minute {
				return fmt.Sprintf("%ds ago", int(d.Seconds()))
			} else if d < time.Hour {
				return fmt.Sprintf("%dm ago", int(d.Minutes()))
			} else if d < 24*time.Hour {
				return fmt.Sprintf("%dh ago", int(d.Hours()))
			}
			return fmt.Sprintf("%dd ago", int(d.Hours()/24))
		},
		"reliabilityClass": func(pct float64) string {
			if pct >= 95 {
				return "excellent"
			} else if pct >= 80 {
				return "good"
			} else if pct >= 50 {
				return "fair"
			}
			return "poor"
		},
		"truncate": func(s string, n int) string {
			if len(s) <= n {
				return s
			}
			return s[:n] + "..."
		},
		"sortIndicator": func(column, currentSort, currentOrder string) string {
			if column != currentSort {
				return ""
			}
			if currentOrder == "desc" {
				return " ▼"
			}
			return " ▲"
		},
	}).Parse(indexTemplate)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(rw, data); err != nil {
		log.Printf("Template execution error: %v", err)
	}
}

func (w *WebServer) handleAPIStats(rw http.ResponseWriter, r *http.Request) {
	stats, err := w.db.GetAllStats()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(stats)
}

func (w *WebServer) handleAPIStatsByType(rw http.ResponseWriter, r *http.Request) {
	typ := r.URL.Path[len("/api/stats/"):]
	if typ == "" {
		w.handleAPIStats(rw, r)
		return
	}

	allStats, err := w.db.GetAllStats()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	var filtered []ResolverStats
	for _, s := range allStats {
		if s.Type == typ {
			filtered = append(filtered, s)
		}
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(filtered)
}

const indexTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Public DNS Servers Monitor</title>
    <style>
        :root {
            --bg: #1a1a2e;
            --surface: #16213e;
            --primary: #0f3460;
            --accent: #e94560;
            --text: #eee;
            --text-dim: #888;
            --excellent: #00c853;
            --good: #64dd17;
            --fair: #ffd600;
            --poor: #ff1744;
        }
        * {
            box-sizing: border-box;
            margin: 0;
            padding: 0;
        }
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, sans-serif;
            background: var(--bg);
            color: var(--text);
            line-height: 1.6;
        }
        .container {
            max-width: 1400px;
            margin: 0 auto;
            padding: 20px;
        }
        header {
            background: var(--surface);
            padding: 20px;
            margin-bottom: 20px;
            border-radius: 8px;
        }
        h1 {
            font-size: 1.8em;
            margin-bottom: 10px;
        }
        .subtitle {
            color: var(--text-dim);
            font-size: 0.9em;
        }
        .filters {
            background: var(--surface);
            padding: 15px 20px;
            border-radius: 8px;
            margin-bottom: 20px;
            display: flex;
            gap: 20px;
            flex-wrap: wrap;
            align-items: center;
        }
        .filter-group {
            display: flex;
            align-items: center;
            gap: 10px;
        }
        .filter-group label {
            color: var(--text-dim);
            font-size: 0.9em;
        }
        select {
            background: var(--primary);
            color: var(--text);
            border: none;
            padding: 8px 12px;
            border-radius: 4px;
            font-size: 0.9em;
            cursor: pointer;
        }
        .stats-summary {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
            gap: 15px;
            margin-bottom: 20px;
        }
        .stat-card {
            background: var(--surface);
            padding: 15px;
            border-radius: 8px;
            text-align: center;
        }
        .stat-value {
            font-size: 1.8em;
            font-weight: bold;
            color: var(--accent);
        }
        .stat-label {
            color: var(--text-dim);
            font-size: 0.85em;
        }
        table {
            width: 100%;
            border-collapse: collapse;
            background: var(--surface);
            border-radius: 8px;
            overflow: hidden;
        }
        th, td {
            padding: 12px 15px;
            text-align: left;
        }
        th {
            background: var(--primary);
            font-weight: 600;
            font-size: 0.85em;
            text-transform: uppercase;
            letter-spacing: 0.5px;
        }
        th a, th span.sortable {
            color: var(--text);
            text-decoration: none;
            cursor: pointer;
        }
        th a:hover, th span.sortable:hover {
            color: var(--accent);
        }
        th .sort-indicator {
            font-size: 0.7em;
            margin-left: 4px;
        }
        tr:nth-child(even) {
            background: rgba(255,255,255,0.02);
        }
        tr:hover {
            background: rgba(255,255,255,0.05);
        }
        .reliability {
            font-weight: bold;
            padding: 4px 8px;
            border-radius: 4px;
            font-size: 0.9em;
        }
        .reliability.excellent { background: var(--excellent); color: #000; }
        .reliability.good { background: var(--good); color: #000; }
        .reliability.fair { background: var(--fair); color: #000; }
        .reliability.poor { background: var(--poor); color: #fff; }
        .type-badge {
            display: inline-block;
            padding: 2px 8px;
            border-radius: 4px;
            font-size: 0.75em;
            font-weight: 600;
            text-transform: uppercase;
            background: var(--primary);
        }
        .resolver-name {
            font-weight: 600;
        }
        .resolver-desc {
            color: var(--text-dim);
            font-size: 0.85em;
            max-width: 300px;
        }
        .rtt {
            font-family: monospace;
        }
        .error-text {
            color: var(--poor);
            font-size: 0.8em;
            max-width: 200px;
            word-break: break-word;
        }
        .timestamp {
            color: var(--text-dim);
            font-size: 0.85em;
        }
        .copy-btn {
            background: transparent;
            border: 1px solid var(--text-dim);
            color: var(--text-dim);
            padding: 2px 6px;
            border-radius: 4px;
            font-size: 0.75em;
            cursor: pointer;
            margin-left: 8px;
            transition: all 0.2s ease;
        }
        .copy-btn:hover {
            border-color: var(--accent);
            color: var(--accent);
        }
        .copy-btn:active {
            transform: scale(0.95);
        }
        .toast {
            position: fixed;
            bottom: 20px;
            right: 20px;
            background: var(--excellent);
            color: #000;
            padding: 12px 20px;
            border-radius: 6px;
            font-weight: 600;
            opacity: 0;
            transform: translateY(20px);
            transition: all 0.3s ease;
            z-index: 1000;
        }
        .toast.show {
            opacity: 1;
            transform: translateY(0);
        }
        footer {
            text-align: center;
            padding: 20px;
            color: var(--text-dim);
            font-size: 0.85em;
        }
        @media (max-width: 768px) {
            .container {
                padding: 10px;
            }
            th, td {
                padding: 8px;
                font-size: 0.85em;
            }
            .resolver-desc {
                display: none;
            }
        }
    </style>
</head>
<body>
    <div class="container">
        <header>
            <h1>Public DNS Servers Monitor</h1>
            <p class="subtitle">Real-time availability and reliability tracking for public encrypted DNS servers and relays</p>
            <p class="subtitle" style="margin-top: 8px;"><a href="https://github.com/DNSCrypt/dnscrypt-resolvers" style="color: var(--accent);">GitHub Repository</a></p>
        </header>

        <div class="filters">
            <div class="filter-group">
                <label for="type-filter">Type:</label>
                <select id="type-filter" onchange="applyFilters()">
                    <option value="">All Types</option>
                    {{range $type, $count := .Types}}
                    <option value="{{$type}}" {{if eq $type $.TypeFilter}}selected{{end}}>{{$type}} ({{$count}})</option>
                    {{end}}
                </select>
            </div>
            <div class="filter-group">
                <label for="sort-by">Sort by:</label>
                <select id="sort-by" onchange="applyFilters()">
                    <option value="">Default (Reliability)</option>
                    <option value="name" {{if eq .SortBy "name"}}selected{{end}}>Name</option>
                    <option value="rtt" {{if eq .SortBy "rtt"}}selected{{end}}>RTT</option>
                    <option value="reliability" {{if eq .SortBy "reliability"}}selected{{end}}>Reliability</option>
                </select>
            </div>
        </div>

        <div class="stats-summary">
            <div class="stat-card">
                <div class="stat-value">{{len .Stats}}</div>
                <div class="stat-label">Total Resolvers</div>
            </div>
            <div class="stat-card">
                <div class="stat-value">{{range $i, $s := .Stats}}{{if ge $s.ReliabilityPct 95.0}}{{end}}{{end}}{{len .Stats}}</div>
                <div class="stat-label">Being Monitored</div>
            </div>
        </div>

        <table>
            <thead>
                <tr>
                    <th><span class="sortable" onclick="sortBy('name')">Name<span class="sort-indicator">{{sortIndicator "name" .SortBy .SortOrder}}</span></span></th>
                    <th><span class="sortable" onclick="sortBy('type')">Type<span class="sort-indicator">{{sortIndicator "type" .SortBy .SortOrder}}</span></span></th>
                    <th><span class="sortable" onclick="sortBy('reliability')">Reliability<span class="sort-indicator">{{sortIndicator "reliability" .SortBy .SortOrder}}</span></span></th>
                    <th>Tests</th>
                    <th><span class="sortable" onclick="sortBy('rtt')">Avg RTT<span class="sort-indicator">{{sortIndicator "rtt" .SortBy .SortOrder}}</span></span></th>
                    <th>Last Success</th>
                    <th>Last Error</th>
                </tr>
            </thead>
            <tbody>
                {{range .Stats}}
                <tr>
                    <td>
                        <div class="resolver-name">{{.Name}}{{if .Stamp}}<button class="copy-btn" onclick="copyStamp('{{.Stamp}}')" title="Copy DNS stamp">Copy</button>{{end}}</div>
                        <div class="resolver-desc">{{truncate .Description 60}}</div>
                    </td>
                    <td><span class="type-badge">{{.Type}}</span></td>
                    <td>
                        {{if gt .TotalTests 0}}
                        <span class="reliability {{reliabilityClass .ReliabilityPct}}">
                            {{printf "%.1f" .ReliabilityPct}}%
                        </span>
                        {{else}}
                        <span class="timestamp">No tests</span>
                        {{end}}
                    </td>
                    <td>{{.SuccessCount}}/{{.TotalTests}}</td>
                    <td class="rtt">
                        {{if gt .AvgRTT 0.0}}{{printf "%.0f" .AvgRTT}}ms{{else}}-{{end}}
                    </td>
                    <td class="timestamp">{{formatTimePtr .LastSuccess}}</td>
                    <td class="error-text">{{truncate .LastError 50}}</td>
                </tr>
                {{end}}
            </tbody>
        </table>

        <footer>
            <p>Last updated: {{formatTime .LastUpdated}}</p>
            <p>Data from <a href="https://github.com/DNSCrypt/dnscrypt-resolvers" style="color: var(--accent)">DNSCrypt/dnscrypt-resolvers</a></p>
        </footer>
    </div>

    <div id="toast" class="toast">DNS stamp copied!</div>

    <script>
        const currentSort = '{{.SortBy}}';
        const currentOrder = '{{.SortOrder}}';
        const currentType = '{{.TypeFilter}}';

        function applyFilters() {
            const type = document.getElementById('type-filter').value;
            const sort = document.getElementById('sort-by').value;
            let url = '?';
            if (type) url += 'type=' + encodeURIComponent(type) + '&';
            if (sort) url += 'sort=' + encodeURIComponent(sort);
            window.location.href = url;
        }

        function sortBy(column) {
            let order = 'asc';
            if (column === currentSort) {
                order = currentOrder === 'asc' ? 'desc' : 'asc';
            }
            let url = '?sort=' + encodeURIComponent(column) + '&order=' + order;
            if (currentType) url += '&type=' + encodeURIComponent(currentType);
            window.location.href = url;
        }

        function copyStamp(stamp) {
            navigator.clipboard.writeText(stamp).then(function() {
                const toast = document.getElementById('toast');
                toast.classList.add('show');
                setTimeout(function() {
                    toast.classList.remove('show');
                }, 2000);
            }).catch(function(err) {
                console.error('Failed to copy:', err);
            });
        }
    </script>
</body>
</html>`

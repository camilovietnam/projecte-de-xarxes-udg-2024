package function

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"status_code"},
	)
)

type httpResponse struct {
	Server  string
	Message string
}

func init() {
	// Register metrics with Prometheus
	fmt.Println("[👉] registered metric: http_requests_total")
	prometheus.MustRegister(httpRequestsTotal)
}

func Handle(w http.ResponseWriter, r *http.Request) {
	// Check the URL path and route the request accordingly
	switch r.URL.Path {
	case "/action":
		handleAction(w, r)
	case "/metrics":
		handleMetrics(w, r)
	default:
		http.NotFound(w, r)
	}
}

func handleMetrics(w http.ResponseWriter, r *http.Request) {
	promhttp.Handler().ServeHTTP(w, r)
}

func handleAction(w http.ResponseWriter, r *http.Request) {
	// Simulate occasional failures (1 out of every 20 requests)
	if rand.Intn(20) == 0 {
		httpRequestsTotal.WithLabelValues(fmt.Sprintf("%d", http.StatusInternalServerError)).Inc()
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	} else {
		httpRequestsTotal.WithLabelValues(fmt.Sprintf("%d", http.StatusOK)).Inc()
	}

	var (
		serverName = os.Getenv("SERVER_NAME")
		lastId     int
	)

	db, err := sql.Open("postgres", "host=postgres port=5432 user=u1062049 password=projecte-de-xarxes dbname=projecte-de-xarxes sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Insert a row into the table
	err = db.QueryRow("INSERT INTO test (name, description) VALUES ($1, $2) RETURNING id", "UNKNOWN", "Description1").Scan(&lastId)
	if err != nil {
		panic(err)
	}

	// Update inserted row with the server name
	_, err = db.Exec("UPDATE test SET name = $1 WHERE id = $2", serverName, lastId)
	if err != nil {
		panic(err)
	}

	// // Delete a row from the table
	// _, err = db.Exec("DELETE FROM test WHERE name = $1", "Row1")
	// if err != nil {
	//     panic(err)
	// }

	_json, _ := json.Marshal(httpResponse{
		Server:  serverName,
		Message: fmt.Sprintf("Database operation successful: %d", lastId),
	})

	w.Header().Add("Content-type", "application/json")
	w.Write(_json)
}

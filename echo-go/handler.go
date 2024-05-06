package function

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type httpResponse struct {
	Server  string
	Message string
}

func Handle(w http.ResponseWriter, r *http.Request) {
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

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	 _ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	var err error
	connStr := "host=postgres user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := db.Ping()
		if err != nil {
			http.Error(w, "db not ready", 500)
			return
		}
		fmt.Fprintln(w, "Hello! DB connected.")
	})

	log.Println("server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

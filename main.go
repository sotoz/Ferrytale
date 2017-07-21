package main

import (
	"net/http"
	"database/sql"
	"log"

	"github.com/sotoz/Ferrytale/controller"
	"github.com/sotoz/Ferrytale/database"
)

type Config struct {
	Host string
}

func main() {
	database.DBCon, err := sql.Open("postgres", "user=myname dbname=dbname sslmode=disable")
	if err != nil {
		log.Fatalf("Could not open database: %s", err)
	}

	c := Config{
		Host: "127.0.0.1:3333",
	}

	http.ListenAndServe(c.Host, controller.Router())
}

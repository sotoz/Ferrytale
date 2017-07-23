package main

import (
	"database/sql"
	"log"
	"net/http"

	"fmt"

	"github.com/sotoz/Ferrytale/controller"
	"github.com/sotoz/Ferrytale/database"
)

type Config struct {
	Host string
}

func main() {

	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@%s(%s)/%s?parseTime=true&time_zone=UTC",
			"root",
			"root",
			"",
			"127.0.0.1:33062",
			"ferrytale",
		),
	)
	if err != nil {
		log.Fatalf("Could not open database: %s", err)
	}
	defer db.Close()

	database.DBCon = db
	c := Config{
		Host: "127.0.0.1:3333",
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Ferrytale started...")
	http.ListenAndServe(c.Host, controller.Router())
}

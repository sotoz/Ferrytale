package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sotoz/ferrytale/controller"
	"github.com/sotoz/ferrytale/database"
)

func main() {

	db, err := sql.Open("mysql", fmt.Sprintf("%s", os.Getenv("DATABASE_URL")))
	if err != nil {
		log.Fatalf("Could not open database: %s", err)
	}
	database.DBCon = db
	err = db.Ping()
	if err != nil {
		log.Fatalf("cannot connect to the database: %s", err)
	}
	defer db.Close()

	log.Print("Ferrytale started...")

	port := ":" + os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(port, controller.Router()))
}

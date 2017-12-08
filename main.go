package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sotoz/Ferrytale/controller"
	"github.com/sotoz/Ferrytale/database"
)

// Config describes the configuration struct for the application.
type Config struct {
	Host string
	Port string
}

func main() {

	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s)/%s?parseTime=true&time_zone=UTC",
			os.Getenv("DATABASE_USER"),
			os.Getenv("DATABASE_PASSWORD"),
			os.Getenv("DATABASE_HOST")+":"+os.Getenv("DATABASE_PORT"),
			"ferrytale",
		),
	)
	if err != nil {
		log.Fatalf("Could not open database: %s", err)
	}
	database.DBCon = db
	err = db.Ping()
	if err != nil {
		log.Fatalf("cannot connect to the database: %s", err)
	}
	defer db.Close()

	c := Config{
		Host: os.Getenv("APPLICATION_HOST"),
		Port: os.Getenv("APPLICATION_PORT"),
	}

	log.Print("Ferrytale started...")

	http.ListenAndServe(c.Host+":"+c.Port, controller.Router())
}

package main

import (
	"database/sql"
	"log"
	"net/http"

	"fmt"

	"github.com/sotoz/Ferrytale/controller"
	"github.com/sotoz/Ferrytale/database"
	"strconv"
	"os"
)

type Config struct {
	Host string
	Port int
}

func main() {

	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@%s(%s)/%s?parseTime=true&time_zone=UTC",
			"root",
			"root",
			"",
			"127.0.0.1:3306",
			"ferrytale",
		),
	)
	if err != nil {
		log.Fatalf("Could not open database: %s", err)
	}
	defer db.Close()

	p, err:= strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("Cannot Parse Environmental variable for port : %s", err)
	}
	c := Config{
		Host: "127.0.0.1",
		Port: p,
	}

	database.DBCon = db
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Ferrytale started...")

	port := strconv.Itoa(c.Port)
	if port == ""{
		log.Fatalf("No Correct Port was given : %s", port)
	}

	url := fmt.Sprintf("%s:%s", c.Host, port)
	if url == ""{
		log.Fatalf("No correct url was given: %s", url)
	}
	http.ListenAndServe(url, controller.Router())
}

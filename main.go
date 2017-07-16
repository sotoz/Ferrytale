package main

import (
	"net/http"
)

type Config struct {
	Host string
}

func main() {
	c := Config{
		Host: "127.0.0.1:3333",
	}
	http.ListenAndServe(c.Host, Router())
}

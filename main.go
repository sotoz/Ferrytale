package main

import (
	"net/http"

	"github.com/sotoz/Ferrytale/controller"
)

type Config struct {
	Host string
}

func main() {
	c := Config{
		Host: "127.0.0.1:3333",
	}

	http.ListenAndServe(c.Host, controller.Router())
}

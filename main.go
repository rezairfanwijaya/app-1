package main

import (
	"log"
	"net/http"

	"github.com/rezairfanwijaya/app-1.git/handler"
)

func main() {
	http.HandleFunc("/cars", handler.Car)
	http.HandleFunc("/users", handler.User)

	if err := http.ListenAndServe(":4545", nil); err != nil {
		log.Fatalf("cant serve server on port 4545, err: %s", err)
	}
}

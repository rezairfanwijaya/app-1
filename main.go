package main

import (
	"log"
	"net/http"

	"github.com/rezairfanwijaya/app-1.git/handler"
)

func main() {
	http.HandleFunc("/cars", handler.GetCarList)
	http.HandleFunc("/users", handler.GetUserList)

	if err := http.ListenAndServe(":4545", nil); err != nil {
		log.Fatalf("failed serve server on port 4545, err: %s", err)
	}
}

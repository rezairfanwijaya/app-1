package main

import (
	"log"
	"net/http"

	"github.com/rezairfanwijaya/app-1.git/handler"
	"github.com/rezairfanwijaya/app-1.git/response"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		resp := response.Success{Data: "pong app-1"}
		res, _ := resp.ToJSON()
		_, _ = w.Write(res)
	})
	http.HandleFunc("/cars", handler.GetCarList)
	// http.HandleFunc("/users", handler.GetUserList)

	if err := http.ListenAndServe(":4545", nil); err != nil {
		log.Fatalf("failed serve server on port 4545, err: %s", err)
	}
}

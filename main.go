package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rezairfanwijaya/app-1.git/handler"
	"github.com/rezairfanwijaya/app-1.git/metric"
	"github.com/rezairfanwijaya/app-1.git/response"
)

func init() {
	_ = prometheus.Register(metric.HTTPRequestTotal)
}

func prometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := mux.CurrentRoute(r)
		path, _ := route.GetPathTemplate()

		metric.HTTPRequestTotal.WithLabelValues(path).Inc()
	})
}

func main() {
	router := mux.NewRouter()
	router.Use(prometheusMiddleware)

	router.Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		resp := response.Success{Data: "pong app-1"}
		res, _ := resp.ToJSON()
		_, _ = w.Write(res)
	})

	router.Path("/cars").HandlerFunc(handler.GetCarList)
	router.Path("/users").HandlerFunc(handler.GetUserList)

	if err := http.ListenAndServe(":4545", router); err != nil {
		log.Fatalf("failed serve server on port 4545, err: %s", err)
	}
}

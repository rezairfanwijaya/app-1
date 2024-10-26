package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rezairfanwijaya/app-1.git/handler"
	"github.com/rezairfanwijaya/app-1.git/metric"
	"github.com/rezairfanwijaya/app-1.git/response"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

func init() {
	_ = prometheus.Register(metric.HTTPRequestTotal)
}

func prometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := NewResponseWriter(w)
		next.ServeHTTP(rw, r)

		route := mux.CurrentRoute(r)
		path, _ := route.GetPathTemplate()

		metric.HTTPRequestTotal.WithLabelValues(path).Inc()
	})
}

func main() {
	router := mux.NewRouter()
	router.Path("/metrics").Handler(promhttp.Handler())
	router.Use(prometheusMiddleware)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		resp := response.Success{Data: "pong app-1"}
		res, _ := resp.ToJSON()
		_, _ = w.Write(res)
	})

	router.HandleFunc("/cars", handler.GetCarList)
	router.HandleFunc("/users", handler.GetUserList)

	if err := http.ListenAndServe(":4545", router); err != nil {
		log.Fatalf("failed serve server on port 4545, err: %s", err)
	}
}

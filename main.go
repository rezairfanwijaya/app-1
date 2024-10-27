package main

import (
	"bytes"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"time"

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
	_ = prometheus.Register(metric.HTTPResponse)
	_ = prometheus.Register(metric.HTTPDuration)
	_ = prometheus.Register(metric.Uptime)
}

func prometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := NewResponseWriter(w)
		next.ServeHTTP(rw, r)
		statusCode := rw.statusCode

		route := mux.CurrentRoute(r)
		path, _ := route.GetPathTemplate()

		timer := prometheus.NewTimer(metric.HTTPDuration.WithLabelValues(path))
		hr := metric.HTTPResponse.WithLabelValues(
			path,
			strconv.Itoa(statusCode),
			r.Method,
		)
		hr.Observe(float64(time.Since(time.Now())))
		metric.HTTPRequestTotal.WithLabelValues(path).Inc()
		timer.ObserveDuration()
	})
}

func trackUpTime() {
	var hostname bytes.Buffer
	cmd := exec.Command("hostname")
	cmd.Stdout = &hostname
	err := cmd.Run()
	if err != nil {
		log.Fatalf("failed running command to get hostname, err: %s", err)
		return
	}

	timer := prometheus.NewTimer(metric.Uptime.WithLabelValues(hostname.String()))
	ticker := time.NewTicker(3 * time.Second)

	for range ticker.C {
		timer.ObserveDuration()
	}
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

	go trackUpTime()
	if err := http.ListenAndServe(":4545", router); err != nil {
		log.Fatalf("failed serve server on port 4545, err: %s", err)
	}
}

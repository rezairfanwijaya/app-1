package metric

import "github.com/prometheus/client_golang/prometheus"

var HTTPRequestTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_total",
		Help: "this metric will show http request total for endpoint",
	},
	[]string{"path"},
)

var ResponseStatus = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_response_status",
		Help: "this metric will show http response status for each path",
	},
	[]string{"status"},
)

package metric

import "github.com/prometheus/client_golang/prometheus"

var HTTPRequestTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_total",
		Help: "this metric will show http request total for endpoint",
	},
	[]string{"path"},
)

var HTTPResponse = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "http_response_status",
		Help: "this metric will show total http response status based on status code",
	},
	[]string{"path", "status", "method"},
)

var HTTPDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "http_duration",
		Help: "this metric will show total duration server response the request based on path",
	},
	[]string{"path"},
)

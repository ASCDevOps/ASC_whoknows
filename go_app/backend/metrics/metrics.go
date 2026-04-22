package metrics

import "github.com/prometheus/client_golang/prometheus"

var HttpRequestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total HTTP requests",
	},
	[]string{"path", "method"},
)

func Register() {
	prometheus.MustRegister(HttpRequestsTotal)
}

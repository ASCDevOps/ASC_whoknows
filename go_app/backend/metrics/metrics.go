package metrics

import "github.com/prometheus/client_golang/prometheus"

var HttpRequestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total HTTP requests",
	},
	[]string{"path", "method"},
)

var LoginAttemptsTotal = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "whoknows_login_attempts_total",
		Help: "Total number of login attempts",
	},
)

var LoginSuccessTotal = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "whoknows_login_success_total",
		Help: "Total successful logins",
	},
)

var LoginFailureTotal = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "whoknows_login_failure_total",
		Help: "Total failed logins",
	},
)

// ✅ NY – kun succesfulde registreringer
var RegisterSuccessTotal = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "whoknows_register_success_total",
		Help: "Total successful user registrations",
	},
)

func Register() {
	prometheus.MustRegister(
		HttpRequestsTotal,
		LoginAttemptsTotal,
		LoginSuccessTotal,
		LoginFailureTotal,
		RegisterSuccessTotal,
	)
}

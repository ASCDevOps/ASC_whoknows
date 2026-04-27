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

var RegisterAttemptsTotal = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "whoknows_register_attempts_total",
		Help: "Total number of register attempts",
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

func Register() {
	prometheus.MustRegister(
		HttpRequestsTotal,
		LoginAttemptsTotal,
		RegisterAttemptsTotal,
		LoginSuccessTotal,
		LoginFailureTotal,
	)

}

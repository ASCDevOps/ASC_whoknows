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

var RegisterSuccessTotal = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "whoknows_register_success_total",
		Help: "Total successful user registrations",
	},
)

var SearchRequestsTotal = prometheus.NewCounter(

	prometheus.CounterOpts{
		Name: "whoknows_search_requests_total",
		Help: "Total number of search requests",
	},
)

var SearchQueriesTotal = prometheus.NewCounterVec(

	prometheus.CounterOpts{
		Name: "whoknows_search_queries_total",
		Help: "Total search queries",
	},
	[]string{"query"},
)

var SearchNoResultsTotal = prometheus.NewCounter(

	prometheus.CounterOpts{
		Name: "whoknows_search_no_results_total",
		Help: "Total searches with no results",
	},
)

var SearchResultsTotal = prometheus.NewCounter(

	prometheus.CounterOpts{
		Name: "whoknows_search_results_total",
		Help: "Total number of results returned",
	},
)

var SearchErrorsTotal = prometheus.NewCounter(

	prometheus.CounterOpts{
		Name: "whoknows_search_errors_total",
		Help: "Total search errors",
	},
)

var SearchRequestDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "whoknows_search_request_duration_seconds",
		Help:    "Duration of search requests in seconds",
		Buckets: prometheus.DefBuckets,
	},
	[]string{"path", "method"},
)

func Register() {
	prometheus.MustRegister(
		HttpRequestsTotal,
		LoginAttemptsTotal,
		LoginSuccessTotal,
		LoginFailureTotal,
		RegisterSuccessTotal,
		SearchRequestsTotal,
		SearchQueriesTotal,
		SearchNoResultsTotal,
		SearchResultsTotal,
		SearchErrorsTotal,
		SearchRequestDuration,
	)
}

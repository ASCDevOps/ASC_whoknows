package middleware

import (
	"net/http"
	"time"

	"whoknows_backend/metrics"
)

func Metrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		metrics.HttpRequestsTotal.WithLabelValues(r.URL.Path, r.Method).Inc()

		_ = time.Since(start)
	})
}

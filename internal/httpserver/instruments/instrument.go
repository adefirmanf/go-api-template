package instruments

import (
	"go-api-template/pkg/metrics"
	"net/http"
	"time"
)

// Metrics .
type Metrics struct {
	Histogram metrics.Histogram
	Count     metrics.Counter
}

// Observing .
func (m *Metrics) HttpRequestDuration(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		begin := time.Now()
		next.ServeHTTP(w, r)
		m.Histogram.With("method", r.Method, "path", r.URL.Path).Observe(time.Since(begin).Seconds())
	})
}

// NewMetrics .
func NewMetrics(h metrics.Histogram) *Metrics {
	return &Metrics{
		Histogram: h,
	}
}

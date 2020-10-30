package app

import (
	"github.com/adefirmanf/go-api-template/pkg/metrics/prometheus"

	"github.com/adefirmanf/go-api-template/pkg/metrics"

	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

// App holds all services required by application.
// i.e Logger, Metrics, or domain business
type App struct {
	*Metrics
}

// Metrics .
type Metrics struct {
	Histogram metrics.Histogram
	// Counter   metrics.Counter
	// Gauge     metrics.Gauge
}

// Logger .
type Logger struct{}

// MetricsApp .
func buildMetricsApp() *Metrics {
	promeHistogram := prometheus.NewHistogramFrom(
		stdprometheus.HistogramOpts{
			Namespace: "addsvc",
			Name:      "http_request_duration_ns",
			Buckets:   []float64{.25, .5, 1, 2.5, 5, 10},
			Help:      "Request duration in nanoseconds",
		}, []string{"method", "path"})

	return &Metrics{
		Histogram: promeHistogram,
	}
}

// New .
func New() *App {
	return &App{
		Metrics: buildMetricsApp(),
	}
}

package httpserver

import (
	"fmt"
	"go-api-template/pkg/metrics"
	"net/http"
)

// Httpserver .
type Httpserver struct {
	Metrics
}

// Metrics .
type Metrics struct {
	Histogram metrics.Histogram
}

// ListenAndServe .
func (h *Httpserver) ListenAndServe(port int) {
	r := NewRouterApp(h.Metrics.Histogram)
	http.ListenAndServe(fmt.Sprintf(":%v", port), r.init())
}

// NewHTTPServer .
func NewHTTPServer(h Httpserver) *Httpserver {
	return &Httpserver{
		Metrics: Metrics{
			Histogram: h.Metrics.Histogram,
		},
	}
}

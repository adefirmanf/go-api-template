package httpserver

import (
	"net/http"

	"github.com/adefirmanf/go-api-template/internal/httpserver/handlers"
	"github.com/adefirmanf/go-api-template/internal/httpserver/instruments"
	"github.com/adefirmanf/go-api-template/pkg/metrics"

	"github.com/go-chi/chi"
)

// App .
type App struct {
	Histogram metrics.Histogram
}

func (a *App) init() http.Handler {
	r := chi.NewRouter()
	metrics := instruments.NewMetrics(a.Histogram)
	r.Use(metrics.HttpRequestDuration)

	echoMessage := handlers.NewEchoMessageHandler()

	r.Route("/{message}", func(r chi.Router) {
		r.Use(handlers.ContextMessage)
		r.Get("/", echoMessage.GetHelloMessage)
	})

	return r
}

// NewRouterApp .
func NewRouterApp(h metrics.Histogram) *App {
	return &App{
		Histogram: h,
	}
}

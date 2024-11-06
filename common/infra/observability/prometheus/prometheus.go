package prometheus

import (
	"log/slog"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Metrics struct{}

func (m *Metrics) Start() {
	http.Handle("/metrics", promhttp.Handler())

	if err := http.ListenAndServe(":2112", nil); err != nil {
		slog.Error("Failed to start prometheus server: " + err.Error())
	}
}

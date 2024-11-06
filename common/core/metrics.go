package core

import "sendzap-checkout/common/infra/observability/prometheus"

func NewMetrics() {
	metrics := prometheus.Metrics{}

	metrics.Start()
}

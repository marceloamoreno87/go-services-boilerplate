package core

import (
	"sendzap-checkout/common/infra/observability/opentelemetry"
	"go.opentelemetry.io/otel/trace"
)

var OPTL trace.Tracer

func NewObservability() {
	opentelemetry := opentelemetry.OpenTelemetry{}

	optl, err := opentelemetry.Start()
	if err != nil {
		panic(err)
	}

	OPTL = optl
}

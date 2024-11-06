package helpers

import (
	"log/slog"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

const (
	REQUEST_BODY  = "request_body"
	RESPONSE_BODY = "response_body"
)

const (
	BEGIN_ERROR      = "begin_error"
	USE_CASE_ERROR   = "use_case_error"
	VALIDATOR_ERROR  = "validator_error"
	REPOSITORY_ERROR = "repository_error"
	ENTITY_ERROR     = "entity_error"
	MAPPER_ERROR     = "mapper_error"
	COMMIT_ERROR     = "commit_error"
	INFRA_ERROR      = "infra_error"
	OAUTH_ERROR      = "oauth_error"
	WEBHOOK_ERROR    = "webhook_error"
	STRIPE_ERROR     = "stripe_error"
	WHATSMEOW_ERROR  = "whatsmeow_error"
)

func LogInfo(span trace.Span, key string, value string) {
	slog.Info(key, span.SpanContext().TraceID().String(), value)

	span.AddEvent(key, trace.WithAttributes(attribute.Key(key).String(value)))
}

func LogError(span trace.Span, key string, value string) {
	// Registra a mensagem de erro no logger
	slog.Error(key, span.SpanContext().TraceID().String(), value)

	// Adiciona um evento ao span de rastreamento
	span.AddEvent(key, trace.WithAttributes(attribute.Key(key).String(value)))
}

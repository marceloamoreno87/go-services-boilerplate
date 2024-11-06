package helpers

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"go.opentelemetry.io/otel/trace"
)

func DecodeRequestBody(
	r *http.Request,
	v interface{},
) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func ValidateRequest(
	span trace.Span,
	w http.ResponseWriter,
	request interface{},
) bool {
	err := Validate(request)
	if err != nil {
		LogError(span, VALIDATOR_ERROR, err.Error())
		ReturnError(w, *NewBadRequestError(err.Error()), span.SpanContext().TraceID().String())
		return false
	}
	return true
}

func HandleTransaction(
	ctx context.Context,
	span trace.Span,
	w http.ResponseWriter,
	db *sql.DB,
	fn func(tx *sql.Tx) (interface{}, *CustomError),
) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		LogError(span, BEGIN_ERROR, err.Error())
		ReturnError(w, *NewBadRequestError(ErrUnexpected.Error()), span.SpanContext().TraceID().String())
		return
	}

	data, customErr := fn(tx)
	if customErr.Error() != "" {
		_ = tx.Rollback()
		LogError(span, USE_CASE_ERROR, customErr.Error())
		ReturnError(w, *customErr, span.SpanContext().TraceID().String())
		return
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		LogError(span, COMMIT_ERROR, err.Error())
		ReturnError(w, *NewBadRequestError(ErrUnexpected.Error()), span.SpanContext().TraceID().String())
		return
	}

	ReturnSuccess(ctx, w, data)
}

func HandleTransactionNoReturn(
	ctx context.Context,
	span trace.Span,
	w http.ResponseWriter,
	db *sql.DB,
	fn func(tx *sql.Tx) (data interface{}, customErr *CustomError),
) any {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		LogError(span, BEGIN_ERROR, err.Error())
		ReturnError(w, *NewBadRequestError(ErrUnexpected.Error()), span.SpanContext().TraceID().String())
		return err
	}

	data, customErr := fn(tx)
	if err != nil {
		_ = tx.Rollback()
		LogError(span, USE_CASE_ERROR, err.Error())
		ReturnError(w, *customErr, span.SpanContext().TraceID().String())
		return err
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		LogError(span, COMMIT_ERROR, err.Error())
		ReturnError(w, *NewBadRequestError(ErrUnexpected.Error()), span.SpanContext().TraceID().String())
		return err
	}

	return data
}


func HandleTransactionWithPNG(
	ctx context.Context,
	span trace.Span,
	w http.ResponseWriter,
	db *sql.DB,
	fn func(tx *sql.Tx) (interface{}, *CustomError),
) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		LogError(span, BEGIN_ERROR, err.Error())
		ReturnError(w, *NewBadRequestError(ErrUnexpected.Error()), span.SpanContext().TraceID().String())
		return
	}

	data, customErr := fn(tx)
	if customErr.Error() != "" {
		_ = tx.Rollback()
		LogError(span, USE_CASE_ERROR, customErr.Error())
		ReturnError(w, *customErr, span.SpanContext().TraceID().String())
		return
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		LogError(span, COMMIT_ERROR, err.Error())
		ReturnError(w, *NewBadRequestError(ErrUnexpected.Error()), span.SpanContext().TraceID().String())
		return
	}

	ReturnSuccessPNG(ctx, w, data)
}
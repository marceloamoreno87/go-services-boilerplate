package helpers

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"
)

type customSuccess struct {
	Data any `json:"data"`
}

type customError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Timestamp  string `json:"timestamp"`
	TraceID    string `json:"trace_id"`
}

const (
	CONTENT_TYPE     = "Content-Type"
	APPLICATION_JSON = "application/json"
)

func ReturnSuccess(ctx context.Context, w http.ResponseWriter, data any) {
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(customSuccess{Data: data})
}

func ReturnSuccessPNG(ctx context.Context, w http.ResponseWriter, data any) {
	w.Header().Set(CONTENT_TYPE, "image/png")
	w.Header().Set("Content-Disposition", "attachment; filename=qr.png")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data.([]byte))
}

func ReturnError(w http.ResponseWriter, customErr CustomError, traceID string) {
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.WriteHeader(customErr.StatusCode)
	_ = json.NewEncoder(w).Encode(customError{
		Message:    customErr.Message,
		StatusCode: customErr.StatusCode,
		Timestamp:  time.Now().Format(time.RFC3339),
		TraceID:    traceID,
	})
}

func BadRequest(w http.ResponseWriter, err error, traceID string) {
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(w).Encode(customError{
		Message:    err.Error(),
		StatusCode: http.StatusBadRequest,
		Timestamp:  time.Now().Format(time.RFC3339),
		TraceID:    traceID,
	})
}

func Unauthorized(w http.ResponseWriter, err error, traceID string) {
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.WriteHeader(http.StatusUnauthorized)
	_ = json.NewEncoder(w).Encode(customError{
		Message:    err.Error(),
		StatusCode: http.StatusUnauthorized,
		Timestamp:  time.Now().Format(time.RFC3339),
		TraceID:    traceID,
	})
}
func Redirect(w http.ResponseWriter, r *http.Request, url string) {
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func RedirectToOauth2FrontEnd(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, os.Getenv("OAUTH2_REDIRECT_FRONTEND"), http.StatusTemporaryRedirect)
}

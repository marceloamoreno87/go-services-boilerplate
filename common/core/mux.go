package core

import (
	"github.com/go-chi/chi/v5"
)


func NewMux() *chi.Mux {
	return chi.NewRouter()
}
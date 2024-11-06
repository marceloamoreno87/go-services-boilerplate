package core

import "github.com/go-chi/chi/v5"

type IRouter interface {
	GetRoutes(r *chi.Mux) chi.Router
}

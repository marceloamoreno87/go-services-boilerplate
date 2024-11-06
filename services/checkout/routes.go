package internal

import (
	"sendzap-checkout/common/core"
	"sendzap-checkout/services/checkout/modules/finance"

	"github.com/go-chi/chi/v5"
)

type Router struct{}

func (rtr *Router) GetRoutes(r *chi.Mux) chi.Router {

	authMiddleware := core.AuthMiddleware{}
	r.Use(authMiddleware.CheckAuth())

	r.Route("/v1", func(r chi.Router) {

		r.Route("/checkout", func(r chi.Router) {

			// @Example: /v1/checkout/category
			r.Route("/category", func(r chi.Router) {
				categoryModule := finance.CategoryModule{}
				categoryHandler := categoryModule.SetupCategoryHandler()
				r.Post("/", categoryHandler.CreateCategory)
				r.Put("/{id}", categoryHandler.UpdateCategory)
				r.Delete("/{id}", categoryHandler.DeleteCategory)
				r.Get("/{id}", categoryHandler.FindCategory)
				r.Get("/", categoryHandler.FindAllCategories)
			})

			r.Route("/registration", func(r chi.Router) {
				//@TODO: Implement registration routes
			})

			r.Route("/subscription", func(r chi.Router) {
				//@TODO: Implement subscription routes
			})

			r.Route("/backoffice", func(r chi.Router) {
				//@TODO: Implement backoffice routes
			})

		})

	})

	return r
}

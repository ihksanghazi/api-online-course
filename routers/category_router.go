package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/ihksanghazi/api-online-course/controllers"
	"github.com/ihksanghazi/api-online-course/databases"
	"github.com/ihksanghazi/api-online-course/middlewares"
	"github.com/ihksanghazi/api-online-course/services"
)

func CategoryRouter() *chi.Mux {
	r := chi.NewRouter()

	categoryServices := services.NewCategoryService(databases.DB)
	categoryControllers := controllers.NewCategoryController(categoryServices)

	r.Group(func(r chi.Router) {
		r.Use(middlewares.TokenMiddleware)
		r.Get("/", categoryControllers.FindAll)
		r.Get("/{id}", categoryControllers.FindById)
		r.Post("/", categoryControllers.Create)
		r.Put("/{id}", categoryControllers.Update)
		r.Delete("/{id}", categoryControllers.Delete)
	})

	return r
}

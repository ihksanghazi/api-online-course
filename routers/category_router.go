package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/ihksanghazi/api-online-course/controllers"
	"github.com/ihksanghazi/api-online-course/databases"
	"github.com/ihksanghazi/api-online-course/services"
)

func CategoryRouter() *chi.Mux {
	r := chi.NewRouter()

	validate := validator.New()

	categoryServices := services.NewCategoryService(databases.DB)
	categoryControllers := controllers.NewCategoryController(categoryServices, validate)

	// guest
	r.Get("/", categoryControllers.FindAll)
	r.Get("/{id}", categoryControllers.FindById)

	// only admin
	r.Group(func(r chi.Router) {
		// r.Use(middlewares.TokenMiddleware)
		// r.Use(middlewares.OnlyAdminMiddleware)
		r.Post("/", categoryControllers.Create)
		r.Put("/{id}", categoryControllers.Update)
		r.Delete("/{id}", categoryControllers.Delete)
	})

	return r
}

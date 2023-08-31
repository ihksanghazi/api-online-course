package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/ihksanghazi/api-online-course/controllers"
	"github.com/ihksanghazi/api-online-course/databases"
	"github.com/ihksanghazi/api-online-course/middlewares"
	"github.com/ihksanghazi/api-online-course/services"
)

func ClassRouter() *chi.Mux {
	r := chi.NewRouter()

	validate := validator.New()

	classService := services.NewClassService(databases.DB)
	classController := controllers.NewClassController(classService, validate)

	// guest
	r.Group(func(r chi.Router) {
		r.Get("/", classController.GetAll)
		r.Get("/{id}", classController.GetById)
	})

	// user
	r.Group(func(r chi.Router) {
		r.Use(middlewares.TokenMiddleware)
		r.Post("/add", classController.Invite)
	})

	// teacher
	r.Group(func(r chi.Router) {
		r.Use(middlewares.TokenMiddleware)
		r.Use(middlewares.OnlyTeacherAdminMiddleware)
		r.Post("/", classController.Create)
	})

	return r
}

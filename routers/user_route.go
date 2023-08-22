package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/ihksanghazi/api-online-course/controllers"
	"github.com/ihksanghazi/api-online-course/databases"
	"github.com/ihksanghazi/api-online-course/services"
)

func UserRouters() *chi.Mux {
	r := chi.NewRouter()

	validate := validator.New()

	userService := services.NewUserServices(databases.DB)
	userControllers := controllers.NewUserContollers(userService, validate)

	// auth
	r.Post("/register", userControllers.Register)
	r.Post("/login", userControllers.Login)
	r.Get("/token", userControllers.GetToken)
	r.Delete("/logout", userControllers.Logout)

	r.Get("/", userControllers.GetAllUsers)
	r.Get("/{id}", userControllers.GetUserById)
	return r
}

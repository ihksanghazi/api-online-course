package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/ihksanghazi/api-online-course/controllers"
	"github.com/ihksanghazi/api-online-course/services"
)

func UserRouters() *chi.Mux {
	r := chi.NewRouter()

	userService := services.NewUserServices()
	userControllers := controllers.NewUserContollers(userService)

	r.Post("/register", userControllers.Register)

	return r
}

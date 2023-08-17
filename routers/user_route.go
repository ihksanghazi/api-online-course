package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/ihksanghazi/api-online-course/controllers"
	"github.com/ihksanghazi/api-online-course/databases"
	"github.com/ihksanghazi/api-online-course/services"
)

func UserRouters() *chi.Mux {
	r := chi.NewRouter()

	userService := services.NewUserServices(databases.DB)
	userControllers := controllers.NewUserContollers(userService)

	r.Post("/register", userControllers.Register)
	r.Post("/login", userControllers.Login)
	r.Get("/token", userControllers.GetToken)
	r.Post("/logout", userControllers.Logout)
	return r
}

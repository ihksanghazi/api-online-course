package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ihksanghazi/api-online-course/databases"
	"github.com/ihksanghazi/api-online-course/models"
	"github.com/ihksanghazi/api-online-course/routers"
	"github.com/joho/godotenv"
)

func main() {
	// load env
	err := godotenv.Load()
	if err != nil {
		panic("Error Load Env : " + err.Error())
	}

	// connect db
	databases.ConnectDB()
	fmt.Println("Sukses Koneksi")

	//migrations table
	databases.DB.AutoMigrate(
		&models.User{},
		&models.UserQuizResponse{},
		&models.UserClass{},
		&models.UserAnswer{},
		&models.Quiz{},
		&models.Question{},
		&models.Module{},
		&models.Message{},
		&models.Discussion{},
		&models.Category{},
		&models.Class{},
		&models.ClassModule{},
		&models.ChosenAnswer{},
	)

	r := chi.NewRouter()

	r.Mount("/api/category", routers.CategoryRouter())

	http.ListenAndServe(":5000", r)
}

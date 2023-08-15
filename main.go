package main

import (
	"fmt"

	"github.com/ihksanghazi/api-online-course/databases"
	"github.com/ihksanghazi/api-online-course/models"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error Load Env : " + err.Error())
	}

	databases.ConnectDB()

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

	fmt.Println("Sukses Koneksi")
}

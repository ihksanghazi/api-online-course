package controllers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/ihksanghazi/api-online-course/services"
	"github.com/ihksanghazi/api-online-course/utils"
)

type ClassController interface {
	Create(w http.ResponseWriter, r *http.Request)
}

func NewClassController(Class services.ClassService, Validate *validator.Validate) ClassController {
	return &ClassControllerImpl{
		Class:    Class,
		Validate: Validate,
	}
}

type ClassControllerImpl struct {
	Class    services.ClassService
	Validate *validator.Validate
}

func (c *ClassControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	utils.ResponseJSON(w, http.StatusOK, "Testing", nil)
}

package controllers

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/ihksanghazi/api-online-course/models"
	"github.com/ihksanghazi/api-online-course/services"
	"github.com/ihksanghazi/api-online-course/utils"
	"gorm.io/gorm"
)

type ClassController interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	Invite(w http.ResponseWriter, r *http.Request)
}

func NewClassController(Class services.ClassService, Validate *validator.Validate) ClassController {
	return &ClassControllerImpl{
		ClassService: Class,
		Validate:     Validate,
	}
}

type ClassControllerImpl struct {
	ClassService services.ClassService
	Validate     *validator.Validate
}

func (c *ClassControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	var request models.ClassWebRequest
	//bind json req
	if err := utils.ReadJSON(r, &request); err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	//validasi
	if errMessage := utils.Validation(c.Validate, request); len(errMessage) > 0 {
		utils.ResponseError(w, http.StatusInternalServerError, errMessage)
		return
	}

	classResponse, errService := c.ClassService.Create(request)
	if errService != nil {
		utils.ResponseError(w, http.StatusInternalServerError, errService.Error())
		return
	}

	utils.ResponseJSON(w, http.StatusCreated, "Successfullty created Class", classResponse)

}

func (c *ClassControllerImpl) GetAll(w http.ResponseWriter, r *http.Request) {
	classResponse, classError := c.ClassService.GetAll()
	if classError != nil {
		utils.ResponseError(w, http.StatusInternalServerError, classError.Error())
		return
	}

	utils.ResponseJSON(w, http.StatusOK, "Success Get All Classes", classResponse)

}

func (c *ClassControllerImpl) GetById(w http.ResponseWriter, r *http.Request) {
	classId := chi.URLParam(r, "id")

	if classId == "" {
		utils.ResponseError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	responseClass, errClass := c.ClassService.GetById(classId)

	if errors.Is(errClass, gorm.ErrRecordNotFound) {
		utils.ResponseError(w, http.StatusNotFound, errClass.Error())
		return
	}

	if errClass != nil {
		utils.ResponseError(w, http.StatusInternalServerError, errClass.Error())
		return
	}

	utils.ResponseJSON(w, http.StatusOK, "Success Get Class", responseClass)

}

func (c *ClassControllerImpl) Invite(w http.ResponseWriter, r *http.Request) {
	//bind json request
	var request models.UserClassWebRequest
	if errRequest := utils.ReadJSON(r, &request); errRequest != nil {
		utils.ResponseError(w, http.StatusBadRequest, errRequest.Error())
		return
	}

	classResponse, errResponse := c.ClassService.AddClass(request)
	if errResponse != nil {
		utils.ResponseError(w, http.StatusInternalServerError, errResponse.Error())
		return
	}

	utils.ResponseJSON(w, http.StatusOK, "OK", classResponse)
}

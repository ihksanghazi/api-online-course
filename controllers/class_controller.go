package controllers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/ihksanghazi/api-online-course/models"
	"github.com/ihksanghazi/api-online-course/services"
	"github.com/ihksanghazi/api-online-course/utils"
)

type ClassController interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
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

	var response models.ClassWebResponse
	response.Name = classResponse.Name
	response.CreatedByID = classResponse.CreatedByID
	response.CategoryID = classResponse.CategoryID
	response.Description = classResponse.Description
	response.Thumbnail = classResponse.Thumbnail
	response.Trailer = classResponse.Trailer

	utils.ResponseJSON(w, http.StatusCreated, "Successfullty created Class", response)

}

func (c *ClassControllerImpl) GetAll(w http.ResponseWriter, r *http.Request) {

}

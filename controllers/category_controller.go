package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/ihksanghazi/api-online-course/models"
	"github.com/ihksanghazi/api-online-course/services"
	"github.com/ihksanghazi/api-online-course/utils"
)

type CategoryControllers interface {
	FindAll(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type CategoryControllerImpl struct {
	CategoryService services.CategoryService
	Validator       *validator.Validate
}

func NewCategoryController(categoryService services.CategoryService, Validator *validator.Validate) CategoryControllers {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
		Validator:       Validator,
	}
}

func (c *CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	categories, err := c.CategoryService.FindAll()
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponseJSON(w, http.StatusOK, "Successfully Fetching All Categories", categories)
}

func (c *CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	// ambil parameter id
	id := chi.URLParam(r, "id")

	ResponseCategory, errQuery := c.CategoryService.FindById(id)
	if errQuery != nil {
		utils.ResponseError(w, http.StatusInternalServerError, errQuery.Error())
		return
	}

	utils.ResponseJSON(w, http.StatusOK, "Successfully Fetch Category", ResponseCategory)
}

func (c *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	var request models.CategoryRequest

	// ambil request json
	if err := utils.ReadJSON(r, &request); err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// validasi
	if valMessage := utils.Validation(c.Validator, request); len(valMessage) > 0 {
		utils.ResponseError(w, http.StatusInternalServerError, valMessage)
		return
	}

	ResponseCategory, err := c.CategoryService.Create(&request)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponseJSON(w, http.StatusCreated, "Successfully Created Category", ResponseCategory)
}

func (c *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	// ambil params id
	id := chi.URLParam(r, "id")

	var request models.CategoryRequest

	// Ambil Request JSON
	if err := utils.ReadJSON(r, &request); err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// validasi
	if errMessage := utils.Validation(c.Validator, request); len(errMessage) > 0 {
		utils.ResponseError(w, http.StatusInternalServerError, errMessage)
		return
	}

	ResponseCategory, errQuery := c.CategoryService.Update(&request, id)
	if errQuery != nil {
		utils.ResponseError(w, http.StatusInternalServerError, errQuery.Error())
		return
	}

	utils.ResponseJSON(w, http.StatusOK, "Successfully Updated Category", ResponseCategory)
}

func (c *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	errQuery := c.CategoryService.Delete(id)
	if errQuery != nil {
		utils.ResponseError(w, http.StatusInternalServerError, errQuery.Error())
		return
	}

	utils.ResponseJSON(w, http.StatusOK, "Successfully Deleted Category", nil)
}

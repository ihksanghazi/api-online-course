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

	var categories []models.Category

	categories, err := c.CategoryService.FindAll(&categories)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponseJSON(w, http.StatusOK, "successfully fetching all data", categories)
}

func (c *CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	// ambil parameter id
	id := chi.URLParam(r, "id")

	var category models.Category

	ResponseCategory, errQuery := c.CategoryService.FindById(&category, id)
	if errQuery != nil {
		utils.ResponseError(w, http.StatusInternalServerError, errQuery.Error())
		return
	}

	utils.ResponseJSON(w, http.StatusOK, "successfully Fetch Data", ResponseCategory)
}

func (c *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request) {

	var category models.Category

	// ambil request json
	utils.ReadJSON(r, &category)

	// validasi
	if err := c.Validator.Var(&category.Name, "required,min=3"); err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseCategory, err := c.CategoryService.Create(&category)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponseJSON(w, http.StatusCreated, "successfully created data", ResponseCategory)
}

func (c *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	// ambil params id
	id := chi.URLParam(r, "id")

	var category models.Category

	// Ambil Request JSON
	utils.ReadJSON(r, &category)

	// validasi
	if err := c.Validator.Var(&category.Name, "required,min=3"); err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseCategory, errQuery := c.CategoryService.Update(&category, id)
	if errQuery != nil {
		utils.ResponseError(w, http.StatusInternalServerError, errQuery.Error())
		return
	}

	utils.ResponseJSON(w, http.StatusOK, "successfully updated data", ResponseCategory)
}

func (c *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var category models.Category

	categoryResponse, errQuery := c.CategoryService.Delete(&category, id)
	if errQuery != nil {
		utils.ResponseError(w, http.StatusInternalServerError, errQuery.Error())
		return
	}

	utils.ResponseJSON(w, http.StatusOK, "successfully deleted data", categoryResponse)
}

package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
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
}

func NewCategoryController(categoryService services.CategoryService) CategoryControllers {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (c *CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {

	categories, err := c.CategoryService.FindAll()

	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
	}

	utils.ResponseJSON(w, http.StatusOK, "successfully fetching all data", categories)
}

func (c *CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	// convert string to uuid
	result, err := uuid.Parse(id)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
	}

	var categroy models.Category
	categroy.ID = result

	ResponseCategory, errQuery := c.CategoryService.FindById(&categroy)
	if errQuery != nil {
		utils.ResponseError(w, http.StatusInternalServerError, errQuery.Error())
	}

	utils.ResponseJSON(w, http.StatusOK, "successfully Fetch Data", ResponseCategory)
}

func (c *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request) {

	var category models.Category

	// catching json request from body
	utils.ReadJSON(r, &category)

	ResponseCategory, err := c.CategoryService.Create(&category)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
	}

	utils.ResponseJSON(w, http.StatusCreated, "successfully created data", ResponseCategory)
}

func (c *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	// convert string to uuid
	result, err := uuid.Parse(id)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
	}

	var category models.Category

	// catching json request body
	utils.ReadJSON(r, &category)

	category.ID = result

	ResponseCategory, errQuery := c.CategoryService.Update(&category)
	if errQuery != nil {
		utils.ResponseError(w, http.StatusInternalServerError, errQuery.Error())
	}

	utils.ResponseJSON(w, http.StatusOK, "successfully updated data", ResponseCategory)
}

func (c *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	// convert string to uuid
	result, err := uuid.Parse(id)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
	}

	var category models.Category
	category.ID = result

	categoryResponse, errQuery := c.CategoryService.Delete(&category)
	if errQuery != nil {
		utils.ResponseError(w, http.StatusInternalServerError, errQuery.Error())
	}

	utils.ResponseJSON(w, http.StatusOK, "successfully deleted data", categoryResponse)
}

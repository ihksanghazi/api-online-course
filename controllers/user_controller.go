package controllers

import (
	"net/http"

	"github.com/ihksanghazi/api-online-course/models"
	"github.com/ihksanghazi/api-online-course/services"
	"github.com/ihksanghazi/api-online-course/utils"
)

type UserControllers interface {
	Register(w http.ResponseWriter, r *http.Request)
}

type UserControllersImpl struct {
	UserService services.UserServices
}

func NewUserContollers(UserService services.UserServices) UserControllers {
	return &UserControllersImpl{
		UserService: UserService,
	}
}

func (u *UserControllersImpl) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	// read data json from body
	utils.ReadJSON(r, &user)

	utils.ResponseJSON(w, http.StatusOK, "Testing", user)

}

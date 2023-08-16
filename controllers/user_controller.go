package controllers

import (
	"net/http"

	"github.com/ihksanghazi/api-online-course/models"
	"github.com/ihksanghazi/api-online-course/services"
	"github.com/ihksanghazi/api-online-course/utils"
)

type UserControllers interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
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
	// membaca data json
	utils.ReadJSON(r, &user)

	userResponse, err := u.UserService.Register(&user)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponseJSON(w, http.StatusOK, "Testing", userResponse)

}

func (u *UserControllersImpl) Login(w http.ResponseWriter, r *http.Request) {

}

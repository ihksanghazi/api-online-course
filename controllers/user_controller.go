package controllers

import (
	"net/http"
	"time"

	"github.com/ihksanghazi/api-online-course/models"
	"github.com/ihksanghazi/api-online-course/services"
	"github.com/ihksanghazi/api-online-course/utils"
)

type UserControllers interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	GetToken(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
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
	// read body json
	var user models.User
	utils.ReadJSON(r, &user)

	refreshtoken, responseToken, err := u.UserService.Login(&user)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// set refresh token ke dalam cookie
	cookie := http.Cookie{
		Name:     "refresh_token",
		Value:    refreshtoken,
		HttpOnly: true,
		MaxAge:   int(time.Hour * 24),
	}

	http.SetCookie(w, &cookie)

	// mengembalikan response token dalam bentuk json
	utils.ResponseJSON(w, http.StatusOK, "your accessToken", responseToken)

}

func (u *UserControllersImpl) GetToken(w http.ResponseWriter, r *http.Request) {
	cookie, errCookie := r.Cookie("refresh_token")
	if errCookie != nil {
		utils.ResponseError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	accessToken, errGetToken := u.UserService.GetToken(cookie.Value, &models.User{})
	if errGetToken != nil {
		utils.ResponseError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	utils.ResponseJSON(w, http.StatusOK, "Your New Access Token", accessToken)

}

func (u *UserControllersImpl) Logout(w http.ResponseWriter, r *http.Request) {

}

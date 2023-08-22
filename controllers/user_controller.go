package controllers

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/ihksanghazi/api-online-course/models"
	"github.com/ihksanghazi/api-online-course/services"
	"github.com/ihksanghazi/api-online-course/utils"
)

type UserControllers interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	GetToken(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	GetUserById(w http.ResponseWriter, r *http.Request)
}

type UserControllersImpl struct {
	UserService services.UserServices
	Validation  *validator.Validate
}

func NewUserContollers(UserService services.UserServices, Validation *validator.Validate) UserControllers {
	return &UserControllersImpl{
		UserService: UserService,
		Validation:  Validation,
	}
}

func (u *UserControllersImpl) Register(w http.ResponseWriter, r *http.Request) {
	var request models.RegisterRequest
	// binding json req
	if err := utils.ReadJSON(r, &request); err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// validasi
	if errMessage := utils.Validation(u.Validation, request); len(errMessage) > 0 {
		utils.ResponseError(w, http.StatusInternalServerError, errMessage)
		return
	}

	userResponse, err := u.UserService.Register(&request)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var response models.UserResponse
	response.ID = userResponse.ID
	response.Username = userResponse.Username
	response.Email = userResponse.Email
	response.Role = userResponse.Role

	utils.ResponseJSON(w, http.StatusOK, "Successfully Register New User", response)

}

func (u *UserControllersImpl) Login(w http.ResponseWriter, r *http.Request) {
	var request models.LoginRequest
	// binding json req
	if err := utils.ReadJSON(r, &request); err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// validasi
	if errMessage := utils.Validation(u.Validation, request); len(errMessage) > 0 {
		utils.ResponseError(w, http.StatusInternalServerError, errMessage)
		return
	}

	refreshtoken, AccessToken, err := u.UserService.Login(&request)
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

	// mengembalikan Access token dalam bentuk json
	utils.ResponseJSON(w, http.StatusOK, "Your AccessToken", AccessToken)

}

func (u *UserControllersImpl) GetToken(w http.ResponseWriter, r *http.Request) {
	cookie, errCookie := r.Cookie("refresh_token")
	if errCookie != nil {
		utils.ResponseError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	accessToken, errGetToken := u.UserService.GetToken(cookie.Value)
	if errGetToken != nil {
		utils.ResponseError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	utils.ResponseJSON(w, http.StatusOK, "Your New Access Token", accessToken)

}

func (u *UserControllersImpl) Logout(w http.ResponseWriter, r *http.Request) {
	// Menghapus Cookie
	cookie := http.Cookie{
		Name:     "refresh_token",
		HttpOnly: true,
		MaxAge:   -1,
	}

	http.SetCookie(w, &cookie)

	utils.ResponseJSON(w, http.StatusOK, "you are logged out", nil)
}

func (u *UserControllersImpl) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	responseUsers, err := u.UserService.GetAllUsers()
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponseJSON(w, http.StatusOK, "Successfully Get All Users", responseUsers)
}

func (u *UserControllersImpl) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	responseUser, err := u.UserService.GetUserById(id)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponseJSON(w, http.StatusOK, "Successfully Get User", responseUser)
}

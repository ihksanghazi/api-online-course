package services

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ihksanghazi/api-online-course/middlewares"
	"github.com/ihksanghazi/api-online-course/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServices interface {
	Register(request *models.RegisterRequest) (models.User, error)
	Login(request *models.LoginRequest) (string, string, error)
	GetToken(refreshToken string) (string, error)
	GetAllUsers() ([]models.UserResponse, error)
	GetUserById(id string) (models.UserResponse, error)
}

type UserServicesImpl struct {
	DB *gorm.DB
}

func NewUserServices(DB *gorm.DB) UserServices {
	return &UserServicesImpl{
		DB: DB,
	}
}

func (u *UserServicesImpl) Register(request *models.RegisterRequest) (models.User, error) {
	var user models.User
	user.Username = request.Username
	user.Email = request.Email
	user.Password = request.Password
	user.Role = request.Role
	// transaction
	err := u.DB.Transaction(func(tx *gorm.DB) error {
		// Cek apakah alamat email sudah ada dalam Database
		if err := tx.Model(&user).First(&user, "email = ?", user.Email).Error; err == nil {
			// User dengan alamat email sudah ada, kembalikan error
			return errors.New("email already exists")
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			// Terjadi error lain selain ErrRecordNotFound
			return err
		}

		// Hash password menggunakan bcrypt
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		// set password yang sudah terenskripsi
		user.Password = string(hashedPassword)

		// Tidak ada user dengan alamat email yang sama, lanjutkan pembuatan user baru
		if err := tx.Model(&user).Create(&user).Error; err != nil {
			return err
		}

		return nil
	})

	return user, err
}

func (u *UserServicesImpl) Login(request *models.LoginRequest) (string, string, error) {
	var user models.User
	var AccessToken string
	var RefreshToken string

	// mulai transaksi
	err := u.DB.Transaction(func(tx *gorm.DB) error {
		// cari berdasarkan email
		if err := tx.Model(&user).First(&user, "email = ?", request.Email).Error; err != nil {
			return err
		}

		// cek hash password
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
			return errors.New("wrong password")
		}

		// membuat refresh token
		claimsRefreshToken := middlewares.ClaimsToken{
			Id: user.ID.String(),
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			},
		}

		tokenRefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefreshToken)
		resultRefreshToken, err := tokenRefreshToken.SignedString([]byte(os.Getenv("REFRESH_JWT_KEY")))
		if err != nil {
			return err
		}

		// simpan refresh token ke dalam database
		if err := tx.Model(&user).Where("id = ?", user.ID).Update("refresh_token", resultRefreshToken).Error; err != nil {
			return err
		}

		// membuat access token
		claimsAccessToken := middlewares.ClaimsToken{
			Id: user.ID.String(),
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 30)),
			},
		}

		tokenAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsAccessToken)
		resultAccessToken, err := tokenAccessToken.SignedString([]byte(os.Getenv("ACCESS_JWT_KEY")))
		if err != nil {
			return err
		}

		AccessToken = resultAccessToken
		RefreshToken = resultRefreshToken

		// commit
		return nil
	})

	return RefreshToken, AccessToken, err
}

func (u *UserServicesImpl) GetToken(refreshToken string) (string, error) {
	var user models.User
	var errorResult error
	var tokenResult string

	// cek refresh token di database
	if err := u.DB.Model(&user).First(&user, "refresh_token = ?", refreshToken).Error; err != nil {
		errorResult = err
	}

	// verify refresh token
	token, errToken := jwt.ParseWithClaims(refreshToken, &middlewares.ClaimsToken{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("REFRESH_JWT_KEY")), nil
	})
	if errToken != nil {
		errorResult = errToken
	}

	if _, ok := token.Claims.(*middlewares.ClaimsToken); ok && token.Valid {
		// Jika refresh token valid, Buat access token
		claimsAccessToken := middlewares.ClaimsToken{
			Id: user.ID.String(),
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 20)),
			},
		}

		tokenAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsAccessToken)
		resultAccessToken, errSigned := tokenAccessToken.SignedString([]byte(os.Getenv("ACCESS_JWT_KEY")))
		if errSigned != nil {
			errorResult = errSigned
		}
		tokenResult = resultAccessToken
	}

	return tokenResult, errorResult
}

func (u *UserServicesImpl) GetAllUsers() ([]models.UserResponse, error) {
	var user []models.User
	var response []models.UserResponse
	var resultErr error

	if err := u.DB.Model(&user).Find(&response, "role != ?", "admin").Error; err != nil {
		resultErr = err
	}

	return response, resultErr
}

func (u *UserServicesImpl) GetUserById(id string) (models.UserResponse, error) {
	var user models.User
	var response models.UserResponse
	var resultErr error

	if err := u.DB.Model(&user).Preload("Classes").Preload("UserClasses").Find(&response, "id = ? AND role != ?", id, "admin").Error; err != nil {
		resultErr = err
	}

	return response, resultErr
}

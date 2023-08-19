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
	Register(modelUser *models.User) (models.User, error)
	Login(modelUser *models.User) (string, string, error)
	GetToken(refreshToken string, userModel *models.User) (string, error)
	GetAllUsers()
}

type UserServicesImpl struct {
	DB *gorm.DB
}

func NewUserServices(DB *gorm.DB) UserServices {
	return &UserServicesImpl{
		DB: DB,
	}
}

func (u *UserServicesImpl) Register(modelUser *models.User) (models.User, error) {
	// transaction
	err := u.DB.Transaction(func(tx *gorm.DB) error {
		// Cek apakah alamat email sudah ada dalam Database
		if err := tx.Where("email = ?", modelUser.Email).First(&modelUser).Error; err == nil {
			// User dengan alamat email sudah ada, kembalikan error
			return errors.New("email already exists")
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			// Terjadi error lain selain ErrRecordNotFound
			return err
		}

		// Hash password menggunakan bcrypt
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(modelUser.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		// set password yang sudah terenskripsi
		modelUser.Password = string(hashedPassword)

		// Tidak ada user dengan alamat email yang sama, lanjutkan pembuatan user baru
		if err := tx.Create(&modelUser).Error; err != nil {
			return err
		}

		return nil
	})

	return *modelUser, err
}

func (u *UserServicesImpl) Login(modelUser *models.User) (string, string, error) {
	var user models.User
	var AccessToken string
	var RefreshToken string

	// mulai transaksi
	err := u.DB.Transaction(func(tx *gorm.DB) error {
		// cari berdasarkan email
		if err := tx.Where("email = ?", modelUser.Email).First(&user).Error; err != nil {
			return err
		}

		// cek hash password
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(modelUser.Password)); err != nil {
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

func (u *UserServicesImpl) GetToken(refreshToken string, userModel *models.User) (string, error) {

	var errorResult error
	var tokenResult string

	// cek refresh token di database
	if err := u.DB.Where("refresh_token = ?", refreshToken).First(&userModel).Error; err != nil {
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
			Id: userModel.ID.String(),
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

func (u *UserServicesImpl) GetAllUsers() {

}

package services

import (
	"errors"

	"github.com/ihksanghazi/api-online-course/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServices interface {
	Register(modelUser *models.User) (models.User, error)
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

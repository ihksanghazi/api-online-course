package services

type UserServices interface {
}

type UserServicesImpl struct {
}

func NewUserServices() UserServices {
	return &UserServicesImpl{}
}

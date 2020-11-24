package services

import (
	"github.com/fullstacktf/Narrativas-Backend/api/interfaces"
)

type UserService struct {
	interfaces.IUserRepository
}

func (service *UserService) Register(username string, password string, email string) error {

	service
}

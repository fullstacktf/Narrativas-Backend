package interfaces

import (
	"github.com/fullstacktf/Narrativas-Backend/api/models"
)

type IUserRepository interface {
	Register(models.User) error
	Login(models.User) error
}

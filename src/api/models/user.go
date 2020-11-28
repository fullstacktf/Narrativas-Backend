package models

import (
	"errors"

	"github.com/fullstacktf/Narrativas-Backend/common"

	"time"
)

type User struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Username  string `gorm:"type:varchar(50); NOT NULL" json:"username"`
	Password  string `gorm:"type:varchar(255); NOT NULL" json:"password"`
	Email     string `gorm:"type:varchar(50); NOT NULL" json:"email"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "user"
}

func (user *User) Register() error {
	var duplicated User
	common.DB.Where("username = ?", &user.Username).Or("email = ?", &user.Email).Find(&duplicated)

	if duplicated.ID != 0 {
		return errors.New("username or email already exists")
	}

	user.Password, _ = common.HashAndSalt([]byte(user.Password))

	if result := common.DB.Omit("Id").Create(user); result.Error != nil {
		return errors.New("invalid data")
	}
	return nil
}

func (user User) Login() (string, error) {
	var test User
	common.DB.Table("user").Select("id", "username", "password").Where("username = ?", user.Username).Scan(&test)

	if err := common.ComparePasswords(test.Password, []byte(user.Password)); err != nil {
		return "", errors.New("invalid username or password")
	}

	token, err := common.CreateToken(uint64(test.ID))

	if err != nil {
		return "", errors.New("unprocessable entity")
	}

	var newLoggedUser = common.UserAuth{
		ID:    test.ID,
		Token: token,
	}

	common.ActiveTokens = append(common.ActiveTokens, newLoggedUser)
	return token, nil
}

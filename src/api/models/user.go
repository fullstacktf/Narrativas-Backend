package models

import (
	"fmt"

	common "github.com/fullstacktf/Narrativas-Backend/common"

	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User : Structure
type User struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Username  string `gorm:"type:varchar(50); NOT NULL" json:"username" binding:"required"`
	Password  string `gorm:"type:varchar(255); NOT NULL" json:"password" binding:"required"`
	Email     string `gorm:"type:varchar(50); NOT NULL" json:"email" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// hashAndSalt : crypts password
func hashAndSalt(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// comparePasswords : checks if a password is correct
func comparePasswords(hashedPwd string, plainPwd []byte) bool {

	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

// TableName : Database table name map
func (User) TableName() string {
	return "user"
}

// Insert : Inserts values into user table
func (user *User) Insert() bool {
	var duplicated User
	common.DB.Where("username = ?", &user.Username).Or("email = ?", &user.Email).Find(&duplicated)

	if duplicated.ID != 0 {
		return true
	}

	user.Password = hashAndSalt([]byte(user.Password))
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	result := common.DB.Omit("Id").Create(user)
	fmt.Println(result.Error)
	return false
}

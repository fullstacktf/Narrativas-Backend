package models

import (
	"errors"
	"regexp"

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

// Check valid email
var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// isEmailValid : checks if the email provided passes the required structure and length.
func isEmailValid(email string) bool {
	if len(email) <= 3 || len(email) > 50 {
		return false
	}
	return emailRegex.MatchString(email)
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
func (user *User) Insert() (bool, error) {
	var duplicated User
	common.DB.Where("username = ?", &user.Username).Or("email = ?", &user.Email).Find(&duplicated)

	if duplicated.ID != 0 {
		return false, errors.New("username or email already exists")
	}

	if !isEmailValid(user.Email) {
		return false, errors.New("invalid email provided")
	}

	user.Password = hashAndSalt([]byte(user.Password))
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	if result := common.DB.Omit("Id").Create(user); result.Error != nil {
		return false, errors.New("invalid data")
	}
	return true, nil
}

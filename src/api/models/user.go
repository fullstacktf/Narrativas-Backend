package models

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/fullstacktf/Narrativas-Backend/common"
	"github.com/fullstacktf/Narrativas-Backend/constants"

	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User : Structure
type User struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Username  string `gorm:"type:varchar(50); NOT NULL" json:"username"`
	Password  string `gorm:"type:varchar(255); NOT NULL" json:"password"`
	Email     string `gorm:"type:varchar(50); NOT NULL" json:"email"`
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

// Register : Register values into user table
func (user *User) Register() (bool, error) {
	var duplicated User
	common.DB.Where("username = ?", &user.Username).Or("email = ?", &user.Email).Find(&duplicated)

	if duplicated.ID != 0 {
		return false, errors.New("username or email already exists")
	}

	user.Password = hashAndSalt([]byte(user.Password))
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	if result := common.DB.Omit("Id").Create(user); result.Error != nil {
		return false, errors.New("invalid data")
	}
	return true, nil
}

// createToken : generates a JWT that lasts 1 hour
func createToken(userid uint64) (string, error) {
	var err error

	os.Setenv("ACCESS_SECRET", os.Getenv(constants.JWTSecret))
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv(constants.AccessSecret)))
	if err != nil {
		return "", err
	}
	return token, nil
}

// Login :  users can log in this endpoint and receive a JTW
func (user User) Login() (string, error) {
	var test User
	common.DB.Debug().Table("user").Select("id", "username", "password").Where("username = ?", user.Username).Scan(&test)

	if !comparePasswords(test.Password, []byte(user.Password)) {
		return "", errors.New("invalid username or password")
	}

	token, err := createToken(uint64(test.ID))
	if err != nil {
		return "", errors.New("unprocessable entity")
	}
	var newLoggedUser = common.UserAuth{
		ID:    test.ID,
		Token: token,
	}

	print(newLoggedUser.ID)
	common.ActiveTokens = append(common.ActiveTokens, newLoggedUser)
	return token, nil
}

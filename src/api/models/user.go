package models

import (
	"errors"
	"net/http"
	"os"
	"regexp"

	"github.com/dgrijalva/jwt-go"
	common "github.com/fullstacktf/Narrativas-Backend/common"
	"github.com/gin-gonic/gin"

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

// A sample use until BBDD is up
var user = User{
	ID:       1,
	Username: "username",
	Password: "password",
}

// createToken : generates a JWT that lasts 15 min
func createToken(userid uint64) (string, error) {
	var err error

	os.Setenv("ACCESS_SECRET", os.Getenv("ROLLIFY_JWT_SECRET"))
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

// Login :  users can log in this endpoint and receive a JTW
func Login(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	// TO DO check user information is OK in database
	if user.Username != u.Username || user.Password != u.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	// // // // // //

	token, err := createToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
}

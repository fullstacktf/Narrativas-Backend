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

type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Username string `gorm:"type:varchar(50); NOT NULL" json:"username"`
	Password string `gorm:"type:varchar(255); NOT NULL" json:"password"`
	Email    string `gorm:"type:varchar(50); NOT NULL" json:"email"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func hashAndSalt(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) error {

	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return err
	}
	return nil
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

	user.Password = hashAndSalt([]byte(user.Password))

	if result := common.DB.Omit("Id").Create(user); result.Error != nil {
		return errors.New("invalid data")
	}
	return nil
}

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

func (user User) Login() (string, error) {
	var test User
	common.DB.Table("user").Select("id", "username", "password").Where("username = ?", user.Username).Scan(&test)

	if err := comparePasswords(test.Password, []byte(user.Password)); err != nil {
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

	common.ActiveTokens = append(common.ActiveTokens, newLoggedUser)
	return token, nil
}

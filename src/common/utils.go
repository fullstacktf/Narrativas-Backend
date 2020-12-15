package common

import (
	"crypto/rand"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func GenerateUUID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	uuid := fmt.Sprintf("%x", b)
	return uuid, nil
}

func CreateDirs(paths []*string) {
	for _, path := range paths {
		os.MkdirAll(*path, os.ModePerm)
	}
}

func HashAndSalt(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePasswords(hashedPwd string, plainPwd []byte) error {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return err
	}
	return nil
}

func CreateToken(userid uint64) (string, error) {
	var err error

	os.Setenv("ACCESS_SECRET", os.Getenv(JWTSecret))
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv(AccessSecret)))
	if err != nil {
		return "", err
	}
	return token, nil
}

func StringInSlice(item string, list []string) bool {
	for _, listItem := range list {
		if item == listItem {
			return true
		}
	}
	return false
}

func init() {
	defaultPath := "./images/default"
	charactersPath := "./images/characters"
	storiesPath := "./images/stories"

	paths := []*string{&defaultPath, &charactersPath, &storiesPath}
	CreateDirs(paths)
}

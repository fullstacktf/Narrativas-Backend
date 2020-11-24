package common

import (
	"crypto/rand"
	"errors"
	"fmt"
	"os"
)

type UserAuth struct {
	ID    uint
	Token string
}

var ActiveTokens []UserAuth

func IsSignedIn(token string) (uint, error) {
	for _, n := range ActiveTokens {
		if token == n.Token {
			return n.ID, nil
		}
	}
	return 0, errors.New("not logged in")
}

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

func StringInSlice(item string, list []string) bool {
	for _, listItem := range list {
		if item == listItem {
			return true
		}
	}
	return false
}

func Init() {
	defaultPath := "./images/default"
	charactersPath := "./images/characters"
	storiesPath := "./images/stories"

	paths := []*string{&defaultPath, &charactersPath, &storiesPath}
	CreateDirs(paths)
}

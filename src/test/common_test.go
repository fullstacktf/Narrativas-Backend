package test

import (
	"fmt"
	"testing"

	"github.com/fullstacktf/Narrativas-Backend/common"
)

func TestStringInSlice(t *testing.T) {
	list := []string{"2", "3", "8", "1", "3", "4"}

	element := "1"
	if common.StringInSlice(element, list) {
		fmt.Printf("String %s is in slice %s\n", element, list)
	} else {
		t.Errorf("String %s is not in slice %s", element, list)
	}

	element = "10"
	if common.StringInSlice(element, list) {
		t.Errorf("String %s is in slice %s", element, list)
	} else {
		fmt.Printf("String %s is not in slice %s\n", element, list)
	}
}

func TestComparePasswords(t *testing.T) {
	password := "password123"
	hashedPassword := "$2a$10$Xz8K0Aeos411odKB6hGc/eb0Hjv4oovo2bujLxbHw/0P59/JY8..u"

	err := common.ComparePasswords(hashedPassword, []byte(password))

	if err != nil {
		t.Errorf("%s is a hash of %s", hashedPassword, password)
	}
}

package test

import (
	"fmt"
	"testing"

	"github.com/fullstacktf/Narrativas-Backend/common"
)

func TestIsSignedIn(t *testing.T) {

	userId := uint(1)
	validToken := "random-token"

	user := common.UserAuth{
		ID:    userId,
		Token: validToken,
	}

	common.ActiveTokens = append(common.ActiveTokens, user)
	invalidToken := "123456"
	userid, err := common.IsSignedIn(invalidToken)

	if err == nil {
		t.Errorf("There is not user signed in with token %s", validToken)
	}

	userid, err = common.IsSignedIn(validToken)
	if userid != userId || err != nil {
		t.Errorf("User with UserID = %d and token %s is logged in", userId, validToken)
	}
}

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

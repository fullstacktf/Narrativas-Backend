package common

import "errors"

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

package common

import "errors"

// UserAuth another struct
type UserAuth struct {
	ID    uint
	Token string
}

var ActiveTokens []UserAuth

// IsSignedIn : Checks if token is valid
func IsSignedIn(token string) (uint, error) {
	for _, n := range ActiveTokens {
		if token == n.Token {
			return n.ID, nil
		}
	}
	return 0, errors.New("not logged in")
}

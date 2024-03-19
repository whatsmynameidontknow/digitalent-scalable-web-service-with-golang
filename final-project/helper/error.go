package helper

import (
	"errors"
	"fmt"
)

var (
	ErrInternal     = errors.New("there's something wrong. it's our fault, not yours")
	ErrUnauthorized = errors.New("you're not allowed to perform this action")
)

func ErrorDuplicate(constraint string) error {
	var field string
	switch constraint {
	case "users_username_key":
		field = "username"
	case "users_email_key":
		field = "email"
	default:
		return errors.New("duplicate value")
	}
	return fmt.Errorf("user with given %s already exists", field)
}

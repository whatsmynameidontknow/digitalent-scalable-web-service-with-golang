package helper

import (
	"errors"
)

var (
	ErrInternal            = errors.New("there's something wrong. it's our fault, not yours")
	ErrUnauthorized        = errors.New("you're not allowed to perform this action")
	ErrDuplicate           = errors.New("user with given username or email already exists")
	ErrInvalidLogin        = errors.New("invalid email or password")
	ErrUserNotFound        = errors.New("user not found")
	ErrPhotoNotFound       = errors.New("photo with given id not found")
	ErrCommentNotFound     = errors.New("comment with given id not found")
	ErrSocialMediaNotFound = errors.New("social media with given id not found")
	ErrNotLoggedIn         = errors.New("you're not logged in")
	ErrInvalidID           = errors.New("id must be a positive integer")
)

type ResponseError struct {
	err  error
	code int
}

func NewResponseError(err error, code int) error {
	return &ResponseError{
		err:  err,
		code: code,
	}
}

func (r *ResponseError) Error() string {
	return r.err.Error()
}

func (r *ResponseError) Code() int {
	return r.code
}

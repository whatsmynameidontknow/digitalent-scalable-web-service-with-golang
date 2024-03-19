package dto

import (
	"errors"
	"regexp"
	"time"
)

type UserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      uint64 `json:"age"`
}

func (u UserRequest) ValidateCreate() error {
	var errs error

	if u.Email == "" {
		errs = errors.Join(errs, errors.New("email can't be empty"))
	}

	if !isValidEmail(u.Email) {
		errs = errors.Join(errs, errors.New("invalid email format"))
	}

	if u.Username == "" {
		errs = errors.Join(errs, errors.New("username can't be empty"))
	}

	if u.Password == "" {
		errs = errors.Join(errs, errors.New("password can't be empty"))
	}

	if len(u.Password) < 6 {
		errs = errors.Join(errs, errors.New("password must be at least 6 characters"))
	}

	// bcrypt.GenerateFromPassword only accepts at most 72 characters
	if len(u.Password) > 72 {
		errs = errors.Join(errs, errors.New("password must be at most 72 characters"))
	}

	if u.Age < 8 {
		errs = errors.Join(errs, errors.New("age must be at least 8 years old"))
	}

	return errs
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return re.MatchString(email)
}

type UserCreateResponse struct {
	ID       uint64 `json:"id"`
	Age      uint64 `json:"age"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func (u UserRequest) ValidateLogin() error {
	var errs error

	if u.Email == "" {
		errs = errors.Join(errs, errors.New("email can't be empty"))
	}

	if u.Password == "" {
		errs = errors.Join(errs, errors.New("password can't be empty"))
	}

	return errs
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

func (u UserRequest) ValidateUpdate() error {
	var errs error

	if u.Username == "" {
		errs = errors.Join(errs, errors.New("username can't be empty"))
	}

	if u.Password == "" {
		errs = errors.Join(errs, errors.New("password can't be empty"))
	}

	if len(u.Password) < 6 {
		errs = errors.Join(errs, errors.New("password must be at least 6 characters"))
	}

	// bcrypt.GenerateFromPassword only accepts at most 72 characters
	if len(u.Password) > 72 {
		errs = errors.Join(errs, errors.New("password must be at most 72 characters"))
	}

	return errs
}

type UserUpdateResponse struct {
	ID        uint64    `json:"id"`
	Age       uint64    `json:"age"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

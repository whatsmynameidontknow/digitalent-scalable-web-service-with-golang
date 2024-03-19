package dto

import (
	"errors"
	"final-project/helper"
	"fmt"
	"time"
)

type PhotoRequest struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
}

func (p PhotoRequest) ValidateCreate() error {
	var errs error

	if p.Title == "" {
		errs = errors.Join(errs, errors.New("title can't be empty"))
	}

	if p.PhotoURL == "" {
		errs = errors.Join(errs, errors.New("photo_url can't be empty"))
	}

	fmt.Println(helper.IsValidURL(p.PhotoURL), p.PhotoURL)
	if !helper.IsValidURL(p.PhotoURL) {
		errs = errors.Join(errs, errors.New("invalid photo_url format"))
	}

	return errs
}

type PhotoCreateResponse struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    uint64    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type PhotoResponse struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    uint64    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User UserResponse `json:"user"`
}

func (p PhotoRequest) ValidateUpdate() error {
	var errs error

	if p.Title == "" {
		errs = errors.Join(errs, errors.New("title can't be empty"))
	}

	if p.PhotoURL == "" {
		errs = errors.Join(errs, errors.New("photo_url can't be empty"))
	}

	fmt.Println(helper.IsValidURL(p.PhotoURL), p.PhotoURL)
	if !helper.IsValidURL(p.PhotoURL) {
		errs = errors.Join(errs, errors.New("invalid photo_url format"))
	}

	return errs
}

type PhotoUpdateResponse struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    uint64    `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

package dto

import (
	"errors"
	"final-project/helper"
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

	if len(p.Title) > 100 {
		errs = errors.Join(errs, errors.New("title can't be more than 100 characters"))
	}

	if p.PhotoURL == "" {
		errs = errors.Join(errs, errors.New("photo_url can't be empty"))
	}

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

	User User `json:"user"`
}

func (p PhotoRequest) ValidateUpdate() error {
	var errs error

	if p.Title == "" {
		errs = errors.Join(errs, errors.New("title can't be empty"))
	}

	if len(p.Title) > 100 {
		errs = errors.Join(errs, errors.New("title can't be more than 100 characters"))
	}

	if p.PhotoURL == "" {
		errs = errors.Join(errs, errors.New("photo_url can't be empty"))
	}

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

type Photo struct {
	ID       uint64 `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID   uint64 `json:"user_id"`
}

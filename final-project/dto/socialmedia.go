package dto

import (
	"errors"
	"final-project/helper"
	"time"
)

type SocialMediaRequest struct {
	Name string `json:"name"`
	URL  string `json:"social_media_url"`
}

func (s SocialMediaRequest) ValidateCreate() error {
	var errs error

	if s.Name == "" {
		errs = errors.Join(errs, errors.New("name can't be empty"))
	}

	if s.URL == "" {
		errs = errors.Join(errs, errors.New("social media url can't be empty"))
	}

	if !helper.IsValidURL(s.URL) {
		errs = errors.Join(errs, errors.New("invalid social media url format"))
	}

	return errs
}

type SocialMediaCreateResponse struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"social_media_url"`
	UserID    uint64    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type SocialMediaResponse struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"social_media_url"`
	UserID    uint64    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user"`
}

func (s SocialMediaRequest) ValidateUpdate() error {
	var errs error

	if s.Name == "" {
		errs = errors.Join(errs, errors.New("name can't be empty"))
	}

	if s.URL == "" {
		errs = errors.Join(errs, errors.New("social media url can't be empty"))
	}

	if !helper.IsValidURL(s.URL) {
		errs = errors.Join(errs, errors.New("invalid social media url format"))
	}

	return errs
}

type SocialMediaUpdateResponse struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"social_media_url"`
	UserID    uint64    `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

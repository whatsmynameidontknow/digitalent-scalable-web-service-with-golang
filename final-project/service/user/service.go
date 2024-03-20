package userservice

import (
	"context"
	"database/sql"
	"errors"
	"final-project/dto"
	"final-project/helper"
	"final-project/model"
	"final-project/repository"
	"final-project/service"
	"net/http"

	"github.com/lib/pq"
)

type userService struct {
	userRepo repository.UserRepository
}

func New(userRepo repository.UserRepository) service.UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) Create(ctx context.Context, data dto.UserRequest) (dto.UserCreateResponse, error) {
	var (
		resp dto.UserCreateResponse
		user model.User
		err  error
	)

	user.Username = data.Username
	user.Email = data.Email
	user.Age = data.Age

	user.Password, err = helper.HashPassword(data.Password)
	if err != nil {
		return resp, helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}

	user, err = u.userRepo.Create(ctx, user)
	if err != nil {
		pqErr := new(pq.Error)
		if errors.As(err, &pqErr) {
			if pqErr.Code.Name() == "unique_violation" {
				return resp, helper.NewResponseError(helper.ErrDuplicate, http.StatusConflict)
			}
		}
		return resp, helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}

	resp.ID = user.ID
	resp.Username = user.Username
	resp.Email = user.Email
	resp.Age = user.Age

	return resp, nil
}

func (u *userService) Login(ctx context.Context, data dto.UserRequest) (dto.UserLoginResponse, error) {
	var resp dto.UserLoginResponse

	user, err := u.userRepo.FindByEmail(ctx, data.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return resp, helper.NewResponseError(helper.ErrInvalidLogin, http.StatusUnauthorized)
		}
		return resp, helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}

	if !helper.IsValidPassword(user.Password, data.Password) {
		return resp, helper.NewResponseError(helper.ErrInvalidLogin, http.StatusUnauthorized)
	}

	resp.Token, err = helper.GenerateJWT(user.ID)
	if err != nil {
		return resp, helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}

	return resp, nil
}

func (u *userService) Update(ctx context.Context, data dto.UserRequest) (dto.UserUpdateResponse, error) {
	var (
		resp dto.UserUpdateResponse
		user model.User
		err  error
	)

	userID, ok := ctx.Value(helper.UserIDKey).(float64)
	if !ok {
		return resp, helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}
	user.ID = uint64(userID)
	user.Email = data.Email
	user.Password, err = helper.HashPassword(data.Password)
	if err != nil {
		return resp, helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}

	user, err = u.userRepo.Update(ctx, user)
	if err != nil {
		pqErr := new(pq.Error)
		if errors.As(err, &pqErr) {
			if pqErr.Code.Name() == "unique_violation" {
				return resp, helper.NewResponseError(helper.ErrDuplicate, http.StatusConflict)
			}
		}
		if errors.Is(err, sql.ErrNoRows) {
			return resp, helper.NewResponseError(helper.ErrUserNotFound, http.StatusNotFound)
		}
		return resp, helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}

	resp.ID = user.ID
	resp.Email = user.Email
	resp.Username = user.Username
	resp.Age = user.Age
	resp.UpdatedAt = user.UpdatedAt

	return resp, nil
}

func (u *userService) Delete(ctx context.Context) error {
	userID, ok := ctx.Value(helper.UserIDKey).(float64)
	if !ok {
		return helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}

	err := u.userRepo.Delete(ctx, uint64(userID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return helper.NewResponseError(helper.ErrUserNotFound, http.StatusNotFound)
		}
		return helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}

	return nil
}

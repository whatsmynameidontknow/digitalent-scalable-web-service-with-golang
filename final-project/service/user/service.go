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
	"log/slog"
	"net/http"

	"github.com/lib/pq"
)

type userService struct {
	userRepo repository.UserRepository
	logger   *slog.Logger
}

func New(userRepo repository.UserRepository, logger *slog.Logger) service.UserService {
	return &userService{userRepo, logger}
}

func (s *userService) Create(ctx context.Context, data dto.UserRequest) (dto.UserCreateResponse, error) {
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
		s.logger.ErrorContext(ctx, err.Error(), "cause", "helper.HashPassword")
		return resp, helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}

	user, err = s.userRepo.Create(ctx, user)
	if err != nil {
		s.logger.ErrorContext(ctx, err.Error(), "cause", "s.userRepo.Create")
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

func (s *userService) Login(ctx context.Context, data dto.UserRequest) (dto.UserLoginResponse, error) {
	var resp dto.UserLoginResponse

	user, err := s.userRepo.FindByEmail(ctx, data.Email)
	if err != nil {
		s.logger.ErrorContext(ctx, err.Error(), "cause", "s.userRepo.FindByEmail")
		if errors.Is(err, sql.ErrNoRows) {
			return resp, helper.NewResponseError(helper.ErrInvalidLogin, http.StatusUnauthorized)
		}
		return resp, helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}

	if !helper.IsValidPassword(user.Password, data.Password) {
		s.logger.ErrorContext(ctx, "invalid password", "cause", "helper.IsValidPassword")
		return resp, helper.NewResponseError(helper.ErrInvalidLogin, http.StatusUnauthorized)
	}

	resp.Token, err = helper.GenerateJWT(user.ID)
	if err != nil {
		s.logger.ErrorContext(ctx, err.Error(), "cause", "helper.GenerateJWT")
		return resp, helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}

	return resp, nil
}

func (s *userService) Update(ctx context.Context, data dto.UserRequest) (dto.UserUpdateResponse, error) {
	var (
		resp dto.UserUpdateResponse
		user model.User
		err  error
	)

	userID, ok := ctx.Value(helper.UserIDKey).(float64)
	if !ok {
		s.logger.ErrorContext(ctx, "userID is not float64", "cause", "ctx.Value(helper.UserIDKey).(float64)")
		return resp, helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}
	user.ID = uint64(userID)
	user.Email = data.Email
	user.Password, err = helper.HashPassword(data.Password)
	if err != nil {
		s.logger.ErrorContext(ctx, err.Error(), "cause", "helper.HashPassword")
		return resp, helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}

	user, err = s.userRepo.Update(ctx, user)
	if err != nil {
		s.logger.ErrorContext(ctx, err.Error(), "cause", "s.userRepo.Update")
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

func (s *userService) Delete(ctx context.Context) error {
	userID, ok := ctx.Value(helper.UserIDKey).(float64)
	if !ok {
		s.logger.ErrorContext(ctx, "userID is not float64", "cause", "ctx.Value(helper.UserIDKey).(float64)")
		return helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}

	err := s.userRepo.Delete(ctx, uint64(userID))
	if err != nil {
		s.logger.ErrorContext(ctx, err.Error(), "cause", "s.userRepo.Delete")
		if errors.Is(err, sql.ErrNoRows) {
			return helper.NewResponseError(helper.ErrUserNotFound, http.StatusNotFound)
		}
		return helper.NewResponseError(helper.ErrInternal, http.StatusInternalServerError)
	}

	return nil
}

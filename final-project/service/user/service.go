package userservice

import (
	"context"
	"errors"
	"final-project/dto"
	"final-project/helper"
	"final-project/model"
	"final-project/repository"
	"final-project/service"

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
		return resp, err
	}

	user, err = u.userRepo.Create(ctx, user)
	if err != nil {
		pqErr := new(pq.Error)
		if errors.As(err, &pqErr) {
			if pqErr.Code.Name() == "unique_violation" {
				return resp, helper.ErrorDuplicate(pqErr.Constraint)
			}
		}
		return resp, err
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
		return resp, err
	}

	if !helper.IsValidPassword(user.Password, data.Password) {
		return resp, errors.New("invalid password")
	}

	resp.Token, err = helper.GenerateJWT(user.ID)
	if err != nil {
		return resp, err
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
		return resp, helper.ErrInternal
	}
	user.ID = uint(userID)
	user.Email = data.Email
	user.Password, err = helper.HashPassword(data.Password)
	if err != nil {
		return resp, err
	}

	user, err = u.userRepo.Update(ctx, user)
	if err != nil {
		return resp, nil
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
		return helper.ErrInternal
	}

	err := u.userRepo.Delete(ctx, uint(userID))
	if err != nil {
		return err
	}

	return nil
}

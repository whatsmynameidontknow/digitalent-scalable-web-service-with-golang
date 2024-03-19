package service

import (
	"context"
	"final-project/dto"
)

type UserService interface {
	Create(context.Context, dto.UserRequest) (dto.UserCreateResponse, error)
	Login(context.Context, dto.UserRequest) (dto.UserLoginResponse, error)
	Update(context.Context, dto.UserRequest) (dto.UserUpdateResponse, error)
	Delete(context.Context) error
}

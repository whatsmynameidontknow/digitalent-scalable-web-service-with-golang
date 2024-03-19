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

type PhotoService interface {
	Create(context.Context, dto.PhotoRequest) (dto.PhotoCreateResponse, error)
	GetAll(context.Context) ([]dto.PhotoResponse, error)
	Update(context.Context, uint64, dto.PhotoRequest) (dto.PhotoUpdateResponse, error)
	Delete(context.Context, uint64) error
}

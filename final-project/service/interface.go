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
	GetByID(context.Context, uint64) (dto.PhotoResponse, error)
}

type CommentService interface {
	Create(context.Context, dto.CommentRequest) (dto.CommentCreateResponse, error)
	GetAll(context.Context) ([]dto.CommentResponse, error)
	Update(context.Context, uint64, dto.CommentRequest) (dto.CommentUpdateResponse, error)
	Delete(context.Context, uint64) error
	GetByID(context.Context, uint64) (dto.CommentResponse, error)
}

type SocialMediaService interface {
	Create(context.Context, dto.SocialMediaRequest) (dto.SocialMediaCreateResponse, error)
	GetAll(context.Context) ([]dto.SocialMediaResponse, error)
	Update(context.Context, uint64, dto.SocialMediaRequest) (dto.SocialMediaUpdateResponse, error)
	Delete(context.Context, uint64) error
	GetByID(context.Context, uint64) (dto.SocialMediaResponse, error)
}

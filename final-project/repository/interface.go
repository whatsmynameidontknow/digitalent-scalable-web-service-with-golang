package repository

import (
	"context"
	"database/sql"
	"final-project/model"
)

type UserRepository interface {
	Create(context.Context, model.User) (model.User, error)
	FindByEmail(context.Context, string) (model.User, error)
	Update(context.Context, model.User) (model.User, error)
	Delete(context.Context, uint64) error
}

type PhotoRepository interface {
	Create(context.Context, model.Photo) (model.Photo, error)
	FindAll(context.Context) ([]model.Photo, error)
	Update(context.Context, *sql.Tx, model.Photo) (model.Photo, error)
	Delete(context.Context, *sql.Tx, uint64) (uint64, error)
	FindByID(context.Context, uint64) (model.Photo, error)
}

type CommentRepository interface {
	Create(context.Context, model.Comment) (model.Comment, error)
	FindAll(context.Context) ([]model.Comment, error)
	Update(context.Context, *sql.Tx, model.Comment) (model.Comment, error)
	Delete(context.Context, *sql.Tx, uint64) (uint64, error)
}

type SocialMediaRepository interface {
	Create(context.Context, model.SocialMedia) (model.SocialMedia, error)
	FindAll(context.Context) ([]model.SocialMedia, error)
	Update(context.Context, *sql.Tx, model.SocialMedia) (model.SocialMedia, error)
	Delete(context.Context, *sql.Tx, uint64) (uint64, error)
}

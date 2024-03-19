package repository

import (
	"context"
	"final-project/model"
)

type UserRepository interface {
	Create(context.Context, model.User) (model.User, error)
	FindByEmail(context.Context, string) (model.User, error)
	Update(context.Context, model.User) (model.User, error)
	Delete(context.Context, uint) error
}

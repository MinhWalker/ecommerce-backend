package repository

import (
	"context"
	"ecommerce-backend/model"
)

type UserRepo interface {
	// User
	SaveUser(context context.Context, user model.User) (model.User, error)
	SelectUserByEmail(context context.Context, email string) (model.User, error)
}

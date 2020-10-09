package repository

import (
	"context"
	"ecommerce-backend/model"
)

type UserRepo interface {
	// User
	SaveUser(context context.Context, user model.User) (model.User, error)
	CheckLogin(context context.Context, email string) (model.User, error)
	GetUserById(context context.Context, userId string) (model.User, error)
	UpdateUser(context context.Context, user model.User) (model.User, error)
	SelectUsers(context context.Context) ([]model.User, error)

	// Product
}

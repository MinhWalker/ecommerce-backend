package repository

import (
	"context"
	"database/sql"
	"ecommerce-backend/db"
	"ecommerce-backend/exception"
	"ecommerce-backend/model"
	"github.com/labstack/gommon/log"
	"github.com/lib/pq"
	"time"
)

type UserRepoImpl struct {
	sql *db.Sql
}

// NewUserRepo create object working with user logic
func NewUserRepo(sql *db.Sql) UserRepo  {
	return UserRepoImpl{
		sql: sql,
	}
}

func (u UserRepoImpl) SaveUser(context context.Context, user model.User) (model.User, error) {
	statement := `
			INSERT INTO users(user_id, email, phone, password, address, full_name, avatar, role, created_at, updated_at)
			VALUES(:user_id, :email, :phone, :password, :address, :full_name, :avatar, :role, :created_at, :updated_at)
	`

	now := time.Now()
	user.CreatedAt = now
	user.UpdateAt = now

	_, err := u.sql.Db.NamedExecContext(context, statement, user)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return user, exception.UserConflict
			}
		}
		return user, exception.SignUpFail
	}

	return user, nil
}

func (u UserRepoImpl) CheckLogin(context context.Context, email string) (model.User, error) {
	var user = model.User{}

	statement := `SELECT * FROM users WHERE email=$1`
	err := u.sql.Db.GetContext(context, &user, statement, email)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, exception.UserNotFound
		}
		log.Error(err.Error())
		return user, err
	}

	return user, nil
}



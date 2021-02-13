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
func NewUserRepo(sql *db.Sql) UserRepo {
	return UserRepoImpl{
		sql: sql,
	}
}

func (u UserRepoImpl) DeleteUsers(context context.Context, userId string) error {
	result := u.sql.Db.MustExecContext(
		context,
		"DELETE FROM users WHERE user_id = $1", userId)

	_, err := result.RowsAffected()
	if err != nil {
		log.Error(err.Error())
		return exception.UserDeleteFail
	}
	return nil

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

func (u UserRepoImpl) GetUserById(context context.Context, userId string) (model.User, error) {
	var user = model.User{}

	statement := `SELECT * FROM users WHERE user_id=$1`
	err := u.sql.Db.GetContext(context, &user, statement, userId)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, exception.UserNotFound
		}
		log.Error(err.Error())
		return user, err
	}

	return user, nil
}

func (u UserRepoImpl) SelectUsers(context context.Context) ([]model.User, error) {
	var users []model.User

	statement := `SELECT * FROM users ORDER BY created_at DESC`
	err := u.sql.Db.SelectContext(context, &users, statement)	// Select many record

	if err != nil {
		if err == sql.ErrNoRows {
			return users, exception.UserEmpty
		}
		log.Error(err.Error())
		return users, err
	}

	return users, nil
}

func (u UserRepoImpl) UpdateUser(context context.Context, user model.User) (model.User, error) {
	sqlStatement := `
		UPDATE users
		SET 
			full_name  = (CASE WHEN LENGTH(:full_name) = 0 THEN full_name ELSE :full_name END),
			email = (CASE WHEN LENGTH(:email) = 0 THEN email ELSE :email END),
			phone = (CASE WHEN LENGTH(:phone) = 0 THEN phone ELSE :phone END),
			address = (CASE WHEN LENGTH(:address) = 0 THEN address ELSE :address END),
			avatar = (CASE WHEN LENGTH(:avatar) = 0 THEN avatar ELSE :avatar END),
			updated_at 	  = COALESCE (:updated_at, updated_at)
		WHERE user_id    = :user_id
	`

	user.UpdateAt = time.Now()

	result, err := u.sql.Db.NamedExecContext(context, sqlStatement, user)
	if err != nil {
		log.Error(err.Error())
		return user, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		log.Error(err.Error())
		return user, exception.UserNotUpdated
	}
	if count == 0 {
		return user, exception.UserNotUpdated
	}

	return user, nil

}

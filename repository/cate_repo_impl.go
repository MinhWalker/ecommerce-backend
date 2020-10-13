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

type CateRepoImpl struct {
	sql *db.Sql
}

// NewUserRepo create object working with user logic
func NewCateRepo(sql *db.Sql) CateRepo {
	return CateRepoImpl{
		sql: sql,
	}
}

func (c CateRepoImpl) SaveCate(context context.Context, cate model.Cate) (model.Cate, error) {
	statement := `
			INSERT INTO categories(cate_id, cate_name, cate_image, created_at, updated_at)
			VALUES(:cate_id, :cate_name, :cate_image, :created_at, :updated_at)
	`

	now := time.Now()
	cate.CreatedAt = now
	cate.UpdatedAt = now

	_, err := c.sql.Db.NamedExecContext(context, statement, cate)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return cate, exception.CateConflict
			}
		}
		return cate, exception.SignUpFail
	}

	return cate, nil
}

func (c CateRepoImpl) DeleteCate(context context.Context, cateId string) error {
	result := c.sql.Db.MustExecContext(
		context,
		"DELETE FROM bookmarks WHERE cate_id = $1", cateId)

	_, err := result.RowsAffected()
	if err != nil {
		log.Error(err.Error())
		return exception.DelBookmarkFail
	}

	return nil
}

func (c CateRepoImpl) UpdateCate(context context.Context, cate model.Cate) error {
	statement := `UPDATE categories
				  SET cate_name = :cate_name, 
					  cate_image = :cate_image 
				  WHERE cate_id=:cate_id;`

	cate.UpdatedAt = time.Now()
	_, err := c.sql.Db.NamedExecContext(context, statement, cate)
	return err
}

func (c CateRepoImpl) SelectCateById(context context.Context, cateId string) (model.Cate, error) {
	var cate = model.Cate{}

	statement := `SELECT * FROM categories WHERE cate_id=$1`
	err := c.sql.Db.GetContext(context, &cate, statement, cateId)

	if err != nil {
		if err == sql.ErrNoRows {
			return cate, exception.CateNotFound
		}
		log.Error(err.Error())
		return cate, err
	}

	return cate, nil
}

func (c CateRepoImpl) SelectCates(context context.Context) ([]model.Cate, error) {
	var cates []model.Cate

	statement := `SELECT * FROM categories ORDER BY created_at DESC`
	err := c.sql.Db.SelectContext(context, &cates, statement)	// Select many record

	if err != nil {
		if err == sql.ErrNoRows {
			return cates, exception.CateEmpty
		}
		log.Error(err.Error())
		return cates, err
	}

	return cates, nil
}


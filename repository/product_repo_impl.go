package repository

import (
	"context"
	"database/sql"
	"ecommerce-backend/db"
	"ecommerce-backend/exception"
	"ecommerce-backend/model"
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"github.com/lib/pq"
	"time"
)

type ProductRepoImpl struct {
	sql *db.Sql
}

func (p ProductRepoImpl) SaveProduct(context context.Context, product model.Product) (model.Product, error) {
	statement := `
			INSERT INTO products(
							product_id, product_name, product_image, product_des,
							cate_id, collection_id, created_at, updated_at)
			VALUES(:product_id, :product_name, :product_image, :product_des,
							:cate_id, :collection_id, :created_at, :updated_at)
	`
	now := time.Now()
	product.CreatedAt = now
	product.UpdatedAt = now

	_, err := p.sql.Db.NamedExecContext(context, statement, product)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return product, errors.New("Product had exits!")
			}
		}
		return product, errors.New("Create new Product fail!")
	}
	return product, nil
}

func (p ProductRepoImpl) AddProductAttribute(context context.Context, productId string, collectionId string, attributes []model.Attribute) error {
	statement := `
		INSERT INTO attributes(
				attr_id, product_id, collection_id, attr_name, size, 
				price, promotion, quantity, created_at, updated_at)
		VALUES(:attr_id, :product_id, :collection_id, :attr_name, 
				:size, :price, :promotion, :quantity, :created_at, :updated_at)
	`

	tx := p.sql.Db.MustBegin()
	for _, attr := range attributes{
		uuid, _ := uuid.NewUUID()
		attr.AttrId = uuid.String()
		attr.ProductId = productId
		attr.CollectionId = collectionId

		now := time.Now()
		attr.CreatedAt = now
		attr.UpdatedAt = now

		_, err := tx.NamedExecContext(context, statement, attr)
		if err != nil {
			log.Error(err.Error())
			if err, ok := err.(*pq.Error); ok {
				if err.Code.Name() == "unique_violation" {
					tx.Commit()
					return errors.New("Product Attribute had exits!")
				}
			}
			tx.Commit()
			return errors.New("Fall to add new Attribute!")
		}
	}
	tx.Commit()
	return nil
}

func (p ProductRepoImpl) UpdateProduct(context context.Context, product model.Product) error {
	panic("implement me")
}

func (p ProductRepoImpl) SelectProductById(context context.Context, productId string) (model.Product, error) {
	var product = model.Product{}
	var attrs []model.Attribute

	statement := `SELECT * FROM products WHERE product_id=$1`
	err := p.sql.Db.GetContext(context, &product, statement, productId)

	if err != nil {
		if err == sql.ErrNoRows {
			return product, exception.ProductNotFound
		}
		log.Error(err.Error())
		return product, err
	}

	statement = `SELECT * FROM attributes WHERE product_id=$1`
	err = p.sql.Db.SelectContext(context, &attrs, statement, productId)

	if err != nil {
		if err == sql.ErrNoRows {
			return product, errors.New("Product don't have this attribute!")
		}
		log.Error(err.Error())
		return product, err
	}

	product.Attributes = attrs

	return product, nil
}

func (p ProductRepoImpl) SelectProducts(context context.Context) ([]model.Product, error) {
	panic("implement me")
}

func NewProductRepo(sql *db.Sql) ProductRepo {
	return ProductRepoImpl{
		sql: sql,
	}
}
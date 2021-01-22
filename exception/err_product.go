package exception

import "errors"

var (
	// product
	ProductNotFound = errors.New("Product not found!")
	ProductEmpty = errors.New("Product empty!")
	DeleteProductFail = errors.New("Fail to delete product")
	DeleteAttributesFail = errors.New("Fail to delete attributes")
	)
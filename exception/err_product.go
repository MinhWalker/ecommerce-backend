package exception

import "errors"

var (
	// product
	ProductNotFound = errors.New("Product not found!")
	ProductEmpty = errors.New("Product empty!")
)
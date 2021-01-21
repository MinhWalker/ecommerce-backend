package exception

import "errors"

var (
	CateConflict   = errors.New("Category has exist!")
	CateNotFound   = errors.New("Can't not find category")
	CateNotUpdated = errors.New("Can't not update category!")
	CateEmpty 	   = errors.New("Empty categories list")
	CateDeleteFail = errors.New("Delete categories fail!")
)

package exception

import "errors"

var (
	UserConflict = errors.New("User has exist!")
	SignUpFail   = errors.New("Fail to create account!")
	UserNotFound = errors.New("Can't not find user")
	UserNotUpdated = errors.New("Can't not update user!")
)

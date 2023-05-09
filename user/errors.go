package user

import "fmt"

var (
	ErrorUserNotFound      = fmt.Errorf("user not found")
	ErrorUserAlreadyExists = fmt.Errorf("user already exists")
)

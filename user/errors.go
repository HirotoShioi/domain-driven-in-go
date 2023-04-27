package user

import "fmt"

var (
	UserNotFound      = fmt.Errorf("User not found")
	UserAlreadyExists = fmt.Errorf("User already exists")
)

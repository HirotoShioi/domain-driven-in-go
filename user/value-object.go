package user

import (
	"fmt"

	"github.com/asaskevich/govalidator"
)

// 冗長的だが、それでいい
// まずはシンプルにやる
var (
	InvalidUserName     = fmt.Errorf("Username must be more than 3 characters long")
	InvalidEmailAddress = fmt.Errorf("Invalid email address")
)

type UserId struct {
	id string
}

func (u *UserId) Equal(userId UserId) bool {
	return u.id == userId.id
}

func (u *UserId) String() string {
	return u.id
}

func (u *UserId) Value() string {
	return u.id
}

func NewUserId(s string) UserId {
	return UserId{s}
}

type UserName struct {
	username string
}

func (u *UserName) Equal(username UserName) bool {
	return u.username == username.username
}

func (u *UserName) Value() string {
	return u.username
}

func (u *UserName) String() string {
	return u.username
}

func NewUserName(s string) (UserName, error) {
	if len(s) < 3 {
		return UserName{}, InvalidUserName
	}

	return UserName{username: s}, nil
}

type Email struct {
	email string
}

func (e *Email) Equal(email Email) bool {
	return e.email == email.email
}

func (e *Email) Value() string {
	return e.email
}

func (e *Email) String() string {
	return e.email
}

func NewEmail(s string) (Email, error) {
	if !govalidator.IsEmail(s) {
		return Email{}, InvalidEmailAddress
	}
	return Email{email: s}, nil
}

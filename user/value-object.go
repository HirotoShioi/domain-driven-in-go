package user

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

// 冗長的だが、それでいい
// まずはシンプルにやる
// TODO: 値オブジェクトの定義を書く
var (
	InvalidUserName     = fmt.Errorf("Username must be more than 3 characters long")
	InvalidEmailAddress = fmt.Errorf("Invalid email address")
)

// TODO: mysqlで自動生成する
// 大文字だと値変更が可能になってしまう
type UserId struct {
	value string
}

func (u *UserId) Equal(userId UserId) bool {
	return u.value == userId.value
}

func (u *UserId) String() string {
	return u.value
}

func (u *UserId) Value() string {
	return u.value
}

func NewUserId() UserId {
	uuid := uuid.New()
	return UserId{uuid.URN()}
}

type UserName struct {
	value string
}

func (u *UserName) Equal(username UserName) bool {
	return u.value == username.value
}

func (u *UserName) Value() string {
	return u.value
}

func (u *UserName) String() string {
	return u.value
}

// バリデーションを強化する
func NewUserName(s string) (UserName, error) {
	if len(s) < 3 {
		return UserName{}, InvalidUserName
	}

	return UserName{value: s}, nil
}

type Email struct {
	value string
}

func (e *Email) Equal(email Email) bool {
	return e.value == email.value
}

func (e *Email) Value() string {
	return e.value
}

func (e *Email) String() string {
	return e.value
}

func NewEmail(s string) (Email, error) {
	if !govalidator.IsEmail(s) {
		return Email{}, InvalidEmailAddress
	}
	return Email{value: s}, nil
}

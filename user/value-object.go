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
type UserId string

type UserName string

type Email string

// バリデーションを強化する
func NewUserName(s string) (UserName, error) {
	if len(s) < 3 {
		return UserName(""), InvalidUserName
	}

	return UserName(s), nil
}

func NewUserId() UserId {
	uuid := uuid.New()
	return UserId(uuid.URN())
}

func NewEmail(s string) (Email, error) {
	if !govalidator.IsEmail(s) {
		return Email(""), InvalidEmailAddress
	}
	return Email(s), nil
}

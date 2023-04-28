// 値オブジェクト
// - システム固有の値
// - (例):ユーザー名をstringとしてではなく、UserNameとして扱うこと
//
// 値オブジェクトの性質
// - 不変である
// - 交換が可能である
// - 等価性によって比較される
//
// 値オブジェクトを利用するメリット
// - 表現力を増す
// - ドキュメント性が高まる
// - 不正な値を存在させない
// - 誤った代入を防ぐ
// - ロジックの散在を防ぐ
package user

import (
	"fmt"

	"github.com/asaskevich/govalidator"
)

var (
	InvalidUserName     = fmt.Errorf("Username must be more than 3 characters long")
	InvalidEmailAddress = fmt.Errorf("Invalid email address")
)

type UserId int32

type UserName string

type Email string

// TODO: バリデーションを強化する
func NewUserName(s string) (UserName, error) {
	if len(s) < 3 {
		return UserName(""), InvalidUserName
	}

	return UserName(s), nil
}

func NewEmail(s string) (Email, error) {
	if !govalidator.IsEmail(s) {
		return Email(""), InvalidEmailAddress
	}
	return Email(s), nil
}

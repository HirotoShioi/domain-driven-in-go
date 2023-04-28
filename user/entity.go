// エンティティ
// - 属性ではなく、同一性(アイデンティティ)によって識別されるオブジェクト
//
// エンティティの性質
// - 可変である
// - 同じ属性であっても区別される
// - 同一性(アイデンティティ)により区別される
//
// 値オブジェクトと何が違う？
// - 明確なライフサイクルが存在し、そこに連続性があること
// - ライフサイクルとは、そのオブジェクトが生成後に属性の可変、削除等の操作を行われること
package user

type User struct {
	id       UserId
	username UserName
	email    Email
}

func NewUser(id UserId, username UserName, email Email) User {
	return User{id: id, username: username, email: email}
}

func (u *User) UserId() UserId     { return u.id }
func (u *User) UserName() UserName { return u.username }
func (u *User) Email() Email       { return u.email }

package user

// TODO: エンティティの定義を書く

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

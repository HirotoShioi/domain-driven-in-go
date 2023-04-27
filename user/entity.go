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

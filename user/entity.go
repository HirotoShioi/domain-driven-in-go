package user

type User struct {
	id       UserId
	username UserName
	email    Email
}

func NewUser(id UserId, username UserName, email Email) User {
	return User{id: id, username: username, email: email}
}

package user

type UserRepository interface {
	Create(user User) (*User, error)
	Delete(userId UserId) error
	Find(userId UserId) (*User, error)
	UpdateUser(userId UserId, username UserName) (*User, error)
}

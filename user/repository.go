package user

type UserRepository interface {
	Create(createUserDto CreateUserDto) (*User, error)
	Delete(userId UserId) error
	Find(userId UserId) (*User, error)
	UpdateUser(updateUserDto UpdateUserDto) (*User, error)
}

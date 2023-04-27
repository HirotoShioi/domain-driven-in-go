package user

// TODO: DTO(Domain transfer object)の定義を書く

type CreateUserDto struct {
	Username UserName
	Email    Email
}

func NewCreateUserDto(username string, email string) (CreateUserDto, error) {
	u, err := NewUserName(username)
	if err != nil {
		return CreateUserDto{}, err
	}
	e, err := NewEmail(email)
	if err != nil {
		return CreateUserDto{}, err
	}

	return CreateUserDto{u, e}, nil
}

type UpdateUserDto struct {
	UserId   UserId
	UserName UserName
}

func NewUpdateUserDto(userId string, username string) (UpdateUserDto, error) {
	u, err := NewUserName(username)
	if err != nil {
		return UpdateUserDto{}, err
	}
	id := UserId{userId}
	return UpdateUserDto{id, u}, nil
}

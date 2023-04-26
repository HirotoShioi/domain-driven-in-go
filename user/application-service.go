package user

import "fmt"

var (
	UserNotFound      = fmt.Errorf("User not found")
	UserAlreadyExists = fmt.Errorf("User already exists")
)

type ApplicationService struct {
	userRepository UserRepository
	userService    UserService
}

func (a *ApplicationService) Register(username, address string) (*User, error) {
	name, err := NewUserName(username)
	if err != nil {
		return nil, err
	}
	email, err := NewEmail(address)
	if err != nil {
		return nil, err
	}
	userId := NewUserId()
	user := NewUser(userId, name, email)

	userExists, err := a.userService.Exists(userId)
	if userExists {
		return nil, UserAlreadyExists
	}

	result, err := a.userRepository.Create(user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *ApplicationService) Get(id string) (*User, error) {
	userId := UserId{id}
	return a.userRepository.Find(userId)
}

func (a *ApplicationService) Update(id, newName string) (*User, error) {
	userId := UserId{id}
	username, err := NewUserName(newName)
	if err != nil {
		return nil, err
	}

	userExists, err := a.userService.Exists(userId)
	if err != nil {
		return nil, err
	}

	if !userExists {
		return nil, UserNotFound
	}

	return a.userRepository.UpdateUser(userId, username)
}

func (a *ApplicationService) Delete(id string) error {
	userId := UserId{id}
	userExists, err := a.userService.Exists(userId)
	if err != nil {
		return err
	}

	if !userExists {
		return nil
	}

	return a.userRepository.Delete(userId)
}

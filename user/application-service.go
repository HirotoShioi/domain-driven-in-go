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

func (a *ApplicationService) Register(createUserDto CreateUserDto) (*UserData, error) {
	result, err := a.userRepository.Create(createUserDto)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *ApplicationService) Get(id string) (*UserData, error) {
	userId := UserId{id}
	return a.userRepository.Find(userId)
}

func (a *ApplicationService) Update(updateUserDto UpdateUserDto) (*UserData, error) {
	userExists, err := a.userService.Exists(updateUserDto.UserId)
	if err != nil {
		return nil, err
	}

	if !userExists {
		return nil, UserNotFound
	}

	return a.userRepository.Update(updateUserDto)
}

func (a *ApplicationService) Delete(userId UserId) error {
	userExists, err := a.userService.Exists(userId)
	if err != nil {
		return err
	}

	if !userExists {
		return nil
	}

	return a.userRepository.Delete(userId)
}

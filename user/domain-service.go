package user

// TODO: ドメインサービスの定義を書く

type UserService struct {
	userRepository UserRepository
}

func (s *UserService) Exists(userId UserId) (bool, error) {
	user, err := s.userRepository.Find(userId)
	if err != nil {
		return false, err
	}
	return user != nil, nil
}

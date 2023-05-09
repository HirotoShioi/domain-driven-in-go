// ドメインサービス
package user

type UserService struct {
	userRepository UserRepository // リポジトリのインターフェースに依存させる
}

func (s *UserService) Exists(userId UserId) (bool, error) {
	user, err := s.userRepository.Find(userId)
	if err != nil && err != ErrorUserNotFound {
		return false, err
	}
	return user != nil, nil
}

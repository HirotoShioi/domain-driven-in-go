package user

// TODO: アプリケーションサービスの定義を書く
// アプリケーションサービス
// - ユースケースの実現
// - ユースケース
// ユーザーの登録
// ユーザー情報の取得
// ユーザー情報の更新
// ユーザーの削除

type ApplicationService struct {
	userRepository UserRepository // リポジトリのインターフェースに依存させる
	userService    UserService
}

func (a *ApplicationService) Register(createUserDto CreateUserDto) (*UserData, error) {
	result, err := a.userRepository.Create(createUserDto)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *ApplicationService) Get(id int32) (*UserData, error) {
	userId := UserId(id)
	return a.userRepository.Find(userId)
}

func (a *ApplicationService) Update(updateUserDto UpdateUserDto) (*UserData, error) {
	userExists, err := a.userService.Exists(updateUserDto.UserId)
	if err != nil {
		return nil, err
	}

	if !userExists {
		return nil, ErrorUserNotFound
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

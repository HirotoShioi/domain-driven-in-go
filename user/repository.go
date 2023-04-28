// リポジトリ
// データの永続化や再構築（復元）を担う
package user

import (
	"context"
	"database/sql"

	db "github.com/HirotoShioi/domain-design/db/sqlc"
)

type UserData struct {
	Userid   int32  `json:"userid"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// UserRepositoryはリポジトリのインターフェース
// リポジトリに依存するコンポーネントはインターフェースに依存させることによってテストが容易になる。
// - 本番環境ではデータベースを利用し、テスト環境ではIn Memory(map等)を使う
type UserRepository interface {
	Create(createUserDto CreateUserDto) (*UserData, error)
	Delete(userId UserId) error
	Find(userId UserId) (*UserData, error)
	Update(updateUserDto UpdateUserDto) (*UserData, error)
}

// UserDBRepositoryはデータベースを利用するリポジトリ
type UserDBRepository struct {
	queries *db.Queries
	context context.Context
}

func NewDatabaseRepository(dbc db.DBTX) *UserDBRepository {
	queries := db.New(dbc)
	return &UserDBRepository{queries: queries, context: context.Background()}
}

func (r *UserDBRepository) Create(createUserDto CreateUserDto) (*UserData, error) {
	arg := db.CreateUserParams{
		Username: string(createUserDto.Username),
		Email:    string(createUserDto.Email),
	}

	dbUser, err := r.queries.CreateUser(context.Background(), arg)
	if err != nil {
		return nil, err
	}
	return toUserData(dbUser), nil
}

func (r *UserDBRepository) Delete(userId UserId) error {
	return r.queries.DeleteUser(r.context, int32(userId))
}

func (r *UserDBRepository) Find(userId UserId) (*UserData, error) {
	user, err := r.queries.GetUser(r.context, int32(userId))
	if err == sql.ErrNoRows {
		return nil, UserNotFound
	} else if err != nil {
		return nil, err
	}
	return toUserData(user), nil
}

func (r *UserDBRepository) Update(dto UpdateUserDto) (*UserData, error) {
	arg := db.UpdateUserParams{
		Userid:   int32(dto.UserId),
		Username: string(dto.UserId),
	}

	user, err := r.queries.UpdateUser(r.context, arg)
	if err != nil {
		return nil, err
	}

	return toUserData(user), nil
}

func toUserData(u db.User) *UserData {
	return &UserData{
		Username: u.Username,
		Userid:   u.Userid,
		Email:    u.Email,
	}
}

package user

// TODO: リポジトリの定義を書く

import (
	"context"
	"database/sql"

	db "github.com/HirotoShioi/domain-design/db/sqlc"
)

type UserData struct {
	Userid   string `json:"userid"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserRepository interface {
	Create(createUserDto CreateUserDto) (*UserData, error)
	Delete(userId UserId) error
	Find(userId UserId) (*UserData, error)
	Update(updateUserDto UpdateUserDto) (*UserData, error)
}

type UserDbRepository struct {
	queries *db.Queries
	context context.Context
}

func NewDatabaseRepository(dbc db.DBTX) *UserDbRepository {
	queries := db.New(dbc)
	return &UserDbRepository{queries: queries, context: context.Background()}
}

func (r *UserDbRepository) Create(createUserDto CreateUserDto) (*UserData, error) {
	arg := db.CreateUserParams{
		Userid:   string(NewUserId()),
		Username: string(createUserDto.Username),
		Email:    string(createUserDto.Email),
	}

	dbUser, err := r.queries.CreateUser(context.Background(), arg)
	if err != nil {
		return nil, err
	}
	return toUserData(dbUser), nil
}

func (r *UserDbRepository) Delete(userId UserId) error {
	return r.queries.DeleteUser(r.context, string(userId))
}

func (r *UserDbRepository) Find(userId UserId) (*UserData, error) {
	user, err := r.queries.GetUser(r.context, string(userId))
	if err == sql.ErrNoRows {
		return nil, UserNotFound
	} else if err != nil {
		return nil, err
	}
	return toUserData(user), nil
}

func (r *UserDbRepository) Update(dto UpdateUserDto) (*UserData, error) {
	arg := db.UpdateUserParams{
		Userid:   string(dto.UserId),
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

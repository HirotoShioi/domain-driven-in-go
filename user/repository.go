package user

import (
	"context"

	db "github.com/HirotoShioi/domain-design/db/sqlc"
)

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
		Userid:   NewUserId().id,
		Username: createUserDto.Username.Value(),
		Email:    createUserDto.Email.Value(),
	}

	dbUser, err := r.queries.CreateUser(context.Background(), arg)
	if err != nil {
		return nil, err
	}
	return toUserData(dbUser), nil
}

func (r *UserDbRepository) Delete(userId UserId) error {
	return r.queries.DeleteUser(r.context, userId.id)
}

func (r *UserDbRepository) Find(userId UserId) (*UserData, error) {
	user, err := r.queries.GetUser(r.context, userId.id)
	if err != nil {
		return nil, err
	}
	return toUserData(user), nil
}

func (r *UserDbRepository) Update(dto UpdateUserDto) (*UserData, error) {
	arg := db.UpdateUserParams{
		Userid:   dto.UserId.id,
		Username: dto.UserId.id,
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
		Id:       u.Userid,
		Email:    u.Email,
	}
}

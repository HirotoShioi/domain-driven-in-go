package db

import (
	"context"
	"testing"
	"time"

	"github.com/HirotoShioi/domain-design/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username: util.RandomUser(),
		Email:    util.RandomEmail(),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.NotZero(t, user.CreatedAt)
	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Userid)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Userid, user2.Userid)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestDeleteUser(t *testing.T) {
	user := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user.Userid)
	require.NoError(t, err)
	_, err = testQueries.GetUser(context.Background(), user.Userid)
	require.Error(t, err)
}

package db

import "testing"

// TODO: それぞれのクエリのテストを実装する
func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username: "test",
		Email:    "email",
	}
}

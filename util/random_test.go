package util

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/stretchr/testify/require"
)

func makeSlice(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = i
	}
	return s
}

func testMany(n int, f func()) {
	for range makeSlice(100) {
		f()
	}
}

func TestRandomString(t *testing.T) {
	for _, i := range makeSlice(100) {
		res := RandomString(i)
		if i == 0 {
			require.Empty(t, res)
		} else {
			require.NotEmpty(t, res)
			require.Len(t, res, i)
		}
	}
}

func TestRandomInt(t *testing.T) {
	testMany(100, func() {
		res := RandomInt(1, 100)
		require.NotEmpty(t, res)
		require.GreaterOrEqual(t, res, int64(1))
		require.LessOrEqual(t, res, int64(100))
	})
}

func TestRandomUser(t *testing.T) {
	testMany(100, func() {
		res := RandomUser()
		require.NotEmpty(t, res)
		require.True(t, govalidator.IsAlphanumeric(res))
	})
}

func TestRandomEmail(t *testing.T) {
	testMany(100, func() {
		res := RandomEmail()
		require.NotEmpty(t, res)
		require.True(t, govalidator.IsEmail(res))
	})

}

package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := "some-password"
	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = CheckPassword(hashedPassword, password)
	require.NoError(t, err)

	wrongPasswornd := "some-wrong-password"
	err = CheckPassword(hashedPassword, wrongPasswornd)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

}

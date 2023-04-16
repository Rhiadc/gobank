package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner:    "Tom",
		Balance:  100,
		Currency: "USD",
	}

	account, err := CreateAccount(arg)
	require.NoError(t, err)

	require.Equal(t, arg.Currency, account.Currency)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Owner, account.Owner)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

}

func CreateAccount(arg CreateAccountParams) (Account, error) {
	return testQueries.CreateAccount(context.Background(), arg)
}

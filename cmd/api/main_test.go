package api

import (
	"testing"
	"time"

	"github.com/rhiadc/gobank/config"
	db "github.com/rhiadc/gobank/db/sqlc"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, *mockdb.MockStoreInterface) *Server {
	config := config.Environments{
		Token: config.Token{
			TokenSymmetricKey:   "aesdreftgydpolikjhyuipokiujhynde",
			AccessTokenDuration: time.Minute,
		},
	}
	server, err := NewServer(&config, store)
	require.NoError(t, err)

	return server
}

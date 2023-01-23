package api

import (
	"testing"
	"time"

	"github.com/rhiadc/gobank/config"
	mockdb "github.com/rhiadc/gobank/db/mock"
	"github.com/stretchr/testify/require"
)

func NewTestServer(t *testing.T, store *mockdb.MockStoreInterface) *Server {
	config := config.Environments{
		Token: config.Token{
			TokenSynmmetricKey:  "aesdreftgydpolikjhyuipokiujhynde",
			AccessTokenDuration: time.Minute,
		},
	}
	server, err := NewServer(&config, store)
	require.NoError(t, err)

	return server
}

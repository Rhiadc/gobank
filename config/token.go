package config

import "time"

type Token struct {
	TokenSynmmetricKey  string
	AccessTokenDuration time.Duration
}

const (
	defaultTokenSymmetricKey   = "12345678901234567890123456789012"
	defaultAccessTokenDuration = "15m"
)

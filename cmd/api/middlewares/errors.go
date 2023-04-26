package middleware

import "errors"

var (
	authHeaderNotProvided = errors.New("authorization header is not provided")
	invalidHeaderFormat   = errors.New("invalid authorization header format")
	unsuportedAuthType    = errors.New("unsuported authorization type")
)

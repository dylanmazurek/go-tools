package totp

import "errors"

var (
	ErrSecretKeyRequired      = errors.New("secret key cannot be empty")
	ErrSecretKeyInvalidLength = errors.New("secret key must be 16, 24, or 32 bytes long")
	ErrSecretKeyInvalidChars  = errors.New("secret key contains invalid characters")
	ErrSecretKeyDecodingBytes = errors.New("secret key failed to decode bytes")

	ErrUnknown = errors.New("unknown error")
)

package totp

import (
	"strings"
	"time"
)

type options struct {
	secretKey *string
	time      *time.Time
}

type Option func(*options)

func WithSecretKey(secretKey string) Option {
	return func(opts *options) {
		upperSecretKey := strings.ToUpper(secretKey)
		opts.secretKey = &upperSecretKey
	}
}

func WithTime(time time.Time) Option {
	return func(opts *options) {
		opts.time = &time
	}
}

func (opts *options) Validate() error {
	if opts.secretKey == nil {
		return ErrSecretKeyRequired
	}

	return nil
}

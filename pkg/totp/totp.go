package totp

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"regexp"
	"strings"
	"time"
)

type TOTP struct {
	secretKey string
	time      *time.Time
}

func New(opts ...Option) (*TOTP, error) {
	o := options{}

	for _, opt := range opts {
		opt(&o)
	}

	err := o.Validate()
	if err != nil {
		return nil, err
	}

	t := &TOTP{
		secretKey: *o.secretKey,
		time:      o.time,
	}

	if isValid, err := t.ValidateSecretKey(); err != nil {
		return nil, err
	} else if isValid {
		return t, nil
	}

	return nil, ErrUnknown
}

func (t *TOTP) ValidateSecretKey() (bool, error) {
	base32Chars := "^[A-Z2-7]+$"
	regex := regexp.MustCompile(base32Chars)

	charsValid := regex.MatchString(t.secretKey)
	lengthValid := len(t.secretKey)%8 == 0

	if !charsValid {
		return false, ErrSecretKeyInvalidChars
	}

	if !lengthValid {
		return false, ErrSecretKeyInvalidLength
	}

	return true, nil
}

func (t *TOTP) Generate() (int, error) {
	if t.time == nil || t.time.IsZero() {
		now := time.Now()
		t.time = &now
	}

	base32Decoder := base32.StdEncoding.WithPadding(base32.NoPadding)
	secretKey := strings.ToUpper(strings.TrimSpace(t.secretKey))
	secretBytes, err := base32Decoder.DecodeString(secretKey)
	if err != nil {
		return 0, ErrSecretKeyDecodingBytes
	}

	timestamp := *t.time
	timeBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(timeBytes, uint64(timestamp.Unix())/30)

	hash := hmac.New(sha1.New, secretBytes)
	hash.Write(timeBytes)
	h := hash.Sum(nil)

	offset := h[len(h)-1] & 0x0F

	truncatedHash := binary.BigEndian.Uint32(h[offset:]) & 0x7FFFFFFF

	return int(truncatedHash % 1_000_000), nil
}

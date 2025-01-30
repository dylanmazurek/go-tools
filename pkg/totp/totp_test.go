package totp_test

import (
	"testing"
	"time"

	"github.com/dylanmazurek/go-tools/pkg/totp"
)

func TestSecretKey(t *testing.T) {
	tests := []struct {
		name           string
		inputSecretKey string
		inputTime      time.Time

		wantIsValid bool
		wantErr     error
	}{
		{
			name:           "secret key valid",
			inputSecretKey: "BLD56GFS34BH67F6",

			wantIsValid: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			opts := []totp.Option{
				totp.WithSecretKey(test.inputSecretKey),
			}

			totp, err := totp.New(opts...)
			if test.wantErr == nil && err != nil {
				t.Errorf("test: %q: got: %v, wantValid: %v", test.name, err, test.wantIsValid)
				return
			}

			if err != nil && test.wantErr != nil && err.Error() != test.wantErr.Error() {
				t.Errorf("test: %q: got: %v, wantErr: %v", test.name, err, test.wantErr)
				return
			}

			if totp == nil && test.wantErr == nil {
				t.Errorf("test: %q: got: %v, wantValid: %v", test.name, totp, test.wantIsValid)
				return
			}

			if err == nil && test.wantIsValid && test.wantErr == nil {
				return
			}

			// validate secret key

			isValid, err := totp.ValidateSecretKey()
			if test.wantErr == nil && err != nil {
				t.Errorf("test: %q: got: %v, wantValid: %v", test.name, err, test.wantIsValid)
				return
			}

			if test.wantIsValid && !isValid {
				t.Errorf("test: %q: got: %v, wantValid: %v", test.name, isValid, test.wantIsValid)
			}
		})
	}
}

func TestValidateCode(t *testing.T) {
	tests := []struct {
		name           string
		inputSecretKey string
		inputTime      time.Time

		wantErr  *string
		wantCode int
	}{
		{
			name:           "secret key valid, time static",
			inputSecretKey: "BLD56SSS34BHU7F6",
			inputTime:      time.Date(2024, 4, 22, 13, 15, 0, 0, time.UTC),

			wantCode: 943648,
		},
		{
			name:           "secret key valid, different time",
			inputSecretKey: "PPD56GFS66BH67F6",
			inputTime:      time.Date(2024, 1, 25, 4, 15, 0, 0, time.UTC),

			wantCode: 407435,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			opts := []totp.Option{
				totp.WithSecretKey(test.inputSecretKey),
				totp.WithTime(test.inputTime),
			}

			totp, err := totp.New(opts...)
			if test.wantErr == nil && err != nil {
				t.Errorf("test: %q got: %v wantErr: %v", test.name, err, test.wantErr)
			}

			code, err := totp.Generate()
			if test.wantErr == nil && err != nil {
				t.Errorf("test %q got: %v wantErr: %v", test.name, err, test.wantErr)
			}

			if test.wantErr == nil && code != test.wantCode {
				t.Errorf("test: %q got: %v wantValid: %v", test.name, code, test.wantCode)
			}

			// want error, check equal
			if test.wantErr != nil && *test.wantErr != err.Error() && err != nil {
				t.Errorf("test %q: got %v, wantErr %v", test.name, err.Error(), test.wantErr)
			}
		})
	}
}

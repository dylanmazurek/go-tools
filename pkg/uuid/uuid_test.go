package uuid_test

import (
	"errors"
	"testing"

	"github.com/dylanmazurek/go-tools/pkg/uuid"
)

func TestParsing(t *testing.T) {
	tests := []struct {
		name  string
		input string

		wantUUIDStr    string
		wantParsedType string
		wantError      error
	}{
		{
			name:  "pass composed id to uuid",
			input: "spaceship-portfolio-earth-b1365644-f8c4-4241-d759-a4abfb87b2e8",

			wantUUIDStr:    "75394081-aa98-ee69-7341-82c6ffa57c1f",
			wantParsedType: uuid.ParseTypeComposed,
		},
		{
			name:  "pass uuid to uuid",
			input: "d09b1eb7-7ed6-4f1c-7ba1-50138688a644",

			wantUUIDStr:    "d09b1eb7-7ed6-4f1c-7ba1-50138688a644",
			wantParsedType: uuid.ParseTypeUUID,
		},
		{
			name:  "pass raw uuid to uuid",
			input: "b1365644-f8c4-4241-957d-a4abfb87b2e8",

			wantUUIDStr:    "b1365644-f8c4-4241-957d-a4abfb87b2e8",
			wantParsedType: uuid.ParseTypeUUID,
		},
		{
			name:  "error empty uuid string",
			input: "",

			wantError: uuid.ErrEmptyString,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uuidObj, errParse := uuid.Parse(test.input)

			uuidString, errString := uuidObj.String()

			err := errors.Join(errParse, errString)
			if err != nil {
				if test.wantError == nil {
					t.Errorf("test %q: got err %q, want nil", test.input, err)
				} else if !errors.Is(err, test.wantError) {
					t.Errorf("test %q: got %q, want err %q", test.input, err, test.wantError)
				}

				return
			}

			if uuidString != nil && *uuidString != test.wantUUIDStr {
				t.Errorf("test %q: got %q, want %q", test.input, *uuidString, test.wantUUIDStr)
			}

			if uuidObj.ParsedUsing != test.wantParsedType {
				t.Errorf("test %q: got %q, want %q", test.input, uuidObj.ParsedUsing, test.wantParsedType)
			}
		})
	}
}

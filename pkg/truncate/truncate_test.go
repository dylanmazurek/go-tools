package truncate_test

import (
	"testing"

	"github.com/dylanmazurek/go-tools/pkg/truncate"
)

func TestString(t *testing.T) {
	tests := []struct {
		name   string
		text   string
		maxLen int
		want   string
	}{
		{
			name:   "Short text less than maxLen",
			text:   "Hello",
			maxLen: 10,
			want:   "Hello",
		},
		{
			name:   "Exact length",
			text:   "Hello World",
			maxLen: 10,
			want:   "Hello World",
		},
		{
			name:   "Long text with space",
			text:   "Hello bright world",
			maxLen: 10,
			want:   "Hello b...",
		},
		{
			name:   "Long text no space",
			text:   "Hello Bright World",
			maxLen: 10,
			want:   "Hello Br...",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := truncate.String(tt.text, tt.maxLen)
			if got != tt.want {
				t.Errorf("expected %q, got %q", tt.want, got)
			}
		})
	}
}

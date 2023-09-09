package cryptoutil_test

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"git.sr.ht/~jamesponddotco/acopw-go/internal/cryptoutil"
)

func TestRandomIndex(t *testing.T) {
	tests := []struct {
		name       string
		n          int
		input      []byte
		expected   int
		expectErr  bool
		errMessage error
	}{
		{
			name:      "ValidRandomIndex",
			n:         10,
			input:     []byte{0x04, 0x00, 0x00, 0x00},
			expected:  4,
			expectErr: false,
		},
		{
			name:      "ExceedsThreshold",
			n:         10,
			input:     append([]byte{0xFF, 0xFF, 0xFF, 0xFF}, 0x04, 0x00, 0x00, 0x00),
			expectErr: false,
		},
		{
			name:       "InvalidRandReader",
			n:          10,
			input:      []byte{0x01},
			expectErr:  true,
			errMessage: io.ErrUnexpectedEOF,
		},
		{
			name:       "ZeroValueN",
			n:          0,
			input:      []byte{0x01, 0x00, 0x00, 0x00},
			expectErr:  true,
			errMessage: cryptoutil.ErrDivisionByZero,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			randReader := bytes.NewReader(tt.input)
			index, err := cryptoutil.RandomIndex(tt.n, randReader)
			if tt.expectErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				if !errors.Is(err, tt.errMessage) {
					t.Fatalf("expected error %v, got %v", tt.errMessage, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if tt.name == "ExceedsThreshold" {
				if index < 0 || index >= tt.n {
					t.Fatalf("index out of valid range [0, %d), got %d", tt.n, index)
				}
			} else {
				if index != tt.expected {
					t.Fatalf("expected %d, got %d", tt.expected, index)
				}
			}
		})
	}
}

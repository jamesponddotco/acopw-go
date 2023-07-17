package acopw_test

import (
	"bytes"
	"io"
	"regexp"
	"testing"

	"git.sr.ht/~jamesponddotco/acopw-go"
)

func TestUUID_Generate(t *testing.T) {
	uuidPattern := `^[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}$`
	uuidRegex := regexp.MustCompile(uuidPattern)

	tests := []struct {
		name  string
		input io.Reader
		err   bool
	}{
		{
			name:  "error reading random source",
			input: bytes.NewBuffer([]byte{}),
			err:   true,
		},
		{
			name:  "nil random source",
			input: nil,
			err:   false,
		},
		{
			name: "successfully generate random UUID",
			input: bytes.NewBuffer([]byte{
				0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
				0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E, 0x1F,
			}),
			err: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uuid := &acopw.UUID{
				Rand: tt.input,
			}

			got, err := uuid.Generate()
			if (err != nil) != tt.err {
				t.Fatalf("expected error %v, got %v", tt.err, err)
			}

			if tt.err {
				return
			}

			if !uuidRegex.MatchString(got) {
				t.Fatalf("expected UUID to match pattern %s, got %s", uuidPattern, got)
			}
		})
	}
}

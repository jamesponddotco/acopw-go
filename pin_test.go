package acopw_test

import (
	"crypto/rand"
	"testing"

	"git.sr.ht/~jamesponddotco/acopw-go"
)

func TestPINGenerate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		pin      acopw.PIN
		validate func(string) bool
	}{
		{
			name: "DefaultLength",
			pin:  acopw.PIN{},
			validate: func(generated string) bool {
				return len(generated) == acopw.DefaultPINLength
			},
		},
		{
			name: "CustomLength",
			pin: acopw.PIN{
				Rand:   rand.Reader,
				Length: 8,
			},
			validate: func(generated string) bool {
				return len(generated) == 8
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := tt.pin.Generate()
			if !tt.validate(got) {
				t.Errorf("PIN.Generate() = %v, validation failed", got)
			}
		})
	}
}

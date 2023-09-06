package acopw_test

import (
	"crypto/rand"
	"errors"
	"testing"

	"git.sr.ht/~jamesponddotco/acopw-go"
)

func TestPIN_Generate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		validate    func(string) bool
		pin         acopw.PIN
		expectedErr error
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
		{
			name: "BigLength",
			pin: acopw.PIN{
				Rand:   rand.Reader,
				Length: 128,
			},
			validate: func(generated string) bool {
				return len(generated) == 128
			},
		},
		{
			name: "FailingReader",
			pin: acopw.PIN{
				Rand: &failingReader{},
			},
			expectedErr: acopw.ErrRandomPIN,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.pin.Generate()
			if !errors.Is(err, tt.expectedErr) {
				t.Fatalf("PIN.Generate() = %v, want %v", err, tt.expectedErr)
			}

			if tt.expectedErr != nil {
				return
			}

			if !tt.validate(got) {
				t.Errorf("PIN.Generate() = %v, validation failed", got)
			}
		})
	}
}

func FuzzPINGenerate(f *testing.F) {
	f.Fuzz(func(t *testing.T, in int) {
		t.Parallel()

		p := &acopw.PIN{
			Rand:   rand.Reader,
			Length: in,
		}

		got, err := p.Generate()
		if err != nil {
			t.Fatal(err)
		}

		if in > 0 && len(got) != in {
			t.Errorf("PIN.Generate() = %v, want %v", got, in)
		}
	})
}

package acopw_test

import (
	"testing"

	"git.sr.ht/~jamesponddotco/acopw-go"
)

func TestPIN_Generate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		validate func(string) bool
		pin      acopw.PIN
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
				Length: 8,
			},
			validate: func(generated string) bool {
				return len(generated) == 8
			},
		},
		{
			name: "BigLength",
			pin: acopw.PIN{
				Length: 128,
			},
			validate: func(generated string) bool {
				return len(generated) == 128
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := tt.pin.Generate()

			if !tt.validate(got) {
				t.Errorf("Random.Generate() = %v, validation failed", got)
			}
		})
	}
}

func FuzzPINGenerate(f *testing.F) {
	f.Fuzz(func(t *testing.T, in int) {
		t.Parallel()

		p := &acopw.PIN{
			Length: in,
		}

		got := p.Generate()

		if in > 0 && len(got) != in {
			t.Errorf("PIN.Generate() = %v, want %v", got, in)
		}
	})
}

package acopw_test

import (
	"crypto/rand"
	"errors"
	"testing"

	"git.sr.ht/~jamesponddotco/acopw-go"
)

type failingReader struct{}

func (*failingReader) Read(_ []byte) (n int, err error) {
	return 0, errors.New("forced read failure")
}

func TestRandom_Generate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		random   *acopw.Random
		validate func(string) bool
		name     string
	}{
		{
			name:   "DefaultConfiguration",
			random: &acopw.Random{},
			validate: func(generated string) bool {
				return len(generated) == acopw.DefaultRandomLength
			},
		},
		{
			name: "CustomLength",
			random: &acopw.Random{
				Rand:   rand.Reader,
				Length: 12,
			},
			validate: func(generated string) bool {
				return len(generated) == 12
			},
		},
		{
			name: "FailingReader",
			random: &acopw.Random{
				Rand: &failingReader{},
			},
			validate: func(generated string) bool {
				return generated == ""
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := tt.random.Generate()
			if !tt.validate(got) {
				t.Errorf("Random.Generate() = %v, validation failed", got)
			}
		})
	}
}

func FuzzRandomGenerate(f *testing.F) {
	f.Fuzz(func(t *testing.T, in int) {
		t.Parallel()

		r := &acopw.Random{
			Rand:   rand.Reader,
			Length: in,
		}

		got := r.Generate()
		if in > 0 && len(got) != in {
			t.Errorf("Random.Generate() = %v, want %v", got, in)
		}
	})
}

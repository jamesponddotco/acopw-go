package acopw_test

import (
	"crypto/rand"
	"errors"
	"strings"
	"sync"
	"testing"

	"git.sr.ht/~jamesponddotco/acopw-go"
)

type failingReader struct{}

func (*failingReader) Read(_ []byte) (n int, err error) {
	return 0, errors.New("forced read failure")
}

type secondFailingReader struct {
	readCount int
	mu        sync.Mutex
}

func (r *secondFailingReader) Read(p []byte) (n int, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.readCount == 0 {
		r.readCount++

		return rand.Reader.Read(p)
	}

	return 0, errors.New("forced read failure on second call")
}

func TestRandom_Generate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		random      *acopw.Random
		validate    func(string) bool
		expectedErr error
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
			expectedErr: acopw.ErrRandomPassword,
		},
		{
			name: "SecondFailingReader",
			random: &acopw.Random{
				Rand:   &secondFailingReader{},
				Length: acopw.DefaultRandomLength,
			},
			expectedErr: acopw.ErrRandomPassword,
		},
		{
			name: "InvalidCharset",
			random: &acopw.Random{
				ExcludedCharset: []string{acopw.Lowercase, acopw.Uppercase, acopw.Numbers, acopw.Symbols},
			},
			expectedErr: acopw.ErrInvalidCharset,
		},
		{
			name: "CharsetExclusions",
			random: &acopw.Random{
				ExcludedCharset: []string{"a", "1", "@"},
			},
			validate: func(generated string) bool {
				return !strings.ContainsAny(generated, "a1@")
			},
		},
		{
			name: "ZeroLength",
			random: &acopw.Random{
				Rand:   rand.Reader,
				Length: 0,
			},
			validate: func(generated string) bool {
				return len(generated) == acopw.DefaultRandomLength
			},
		},
		{
			name: "HighLength",
			random: &acopw.Random{
				Length: 1024,
			},
			validate: func(generated string) bool {
				return len(generated) == 1024
			},
		},
		{
			name: "LowLength",
			random: &acopw.Random{
				Length: 1,
			},
			validate: func(generated string) bool {
				return len(generated) == 1
			},
		},
		{
			name: "UseLowerOnly",
			random: &acopw.Random{
				UseLower:   true,
				UseUpper:   false,
				UseNumbers: false,
				UseSymbols: false,
			},
			validate: func(generated string) bool {
				return strings.ToLower(generated) == generated
			},
		},
		{
			name: "UseUpperOnly",
			random: &acopw.Random{
				UseLower:   false,
				UseUpper:   true,
				UseNumbers: false,
				UseSymbols: false,
			},
			validate: func(generated string) bool {
				return strings.ToUpper(generated) == generated
			},
		},
		{
			name: "UseNumbersOnly",
			random: &acopw.Random{
				UseLower:   false,
				UseUpper:   false,
				UseNumbers: true,
				UseSymbols: false,
			},
			validate: func(generated string) bool {
				for _, char := range generated {
					if !strings.ContainsRune(acopw.Numbers, char) {
						return false
					}
				}

				return true
			},
		},
		{
			name: "UseSymbolsOnly",
			random: &acopw.Random{
				UseLower:   false,
				UseUpper:   false,
				UseNumbers: false,
				UseSymbols: true,
			},
			validate: func(generated string) bool {
				for _, char := range generated {
					if !strings.ContainsRune(acopw.Symbols, char) {
						return false
					}
				}

				return true
			},
		},
		{
			name: "ExcludeCharacter",
			random: &acopw.Random{
				ExcludedCharset: []string{"z"},
			},
			validate: func(generated string) bool {
				return !strings.Contains(generated, "z")
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.random.Generate()
			if !errors.Is(err, tt.expectedErr) {
				t.Fatalf("Random.Generate() = %v, want %v", got, tt.expectedErr)
			}

			if tt.expectedErr != nil {
				return
			}

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

		got, err := r.Generate()
		if err != nil {
			t.Fatal(err)
		}

		if in > 0 && len(got) != in {
			t.Errorf("Random.Generate() = %v, want %v", got, in)
		}
	})
}

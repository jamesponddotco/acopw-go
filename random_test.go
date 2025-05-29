package acopw_test

import (
	"strings"
	"testing"

	"git.sr.ht/~jamesponddotco/acopw-go"
)

func TestRandom_Generate(t *testing.T) {
	t.Parallel()

	const (
		numbers string = "0123456789"
		symbols string = `!"#$%&'()*+,-./:;<=>?@[\]^_{|}~`
	)

	tests := []struct {
		name        string
		random      *acopw.Random
		validate    func(string) bool
		expectPanic bool
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
				Length: 12,
			},
			validate: func(generated string) bool {
				return len(generated) == 12
			},
		},
		{
			name: "CharsetExclusions",
			random: &acopw.Random{
				ExcludedCharset: []string{
					"a",
					"1",
					"@",
				},
			},
			validate: func(generated string) bool {
				return !strings.ContainsAny(generated, "a1@")
			},
		},
		{
			name: "ZeroLength",
			random: &acopw.Random{
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
					if !strings.ContainsRune(numbers, char) {
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
					if !strings.ContainsRune(symbols, char) {
						return false
					}
				}

				return true
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if tt.expectPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Random.Generate() did not panic")
					}
				}()
			}

			got := tt.random.Generate()

			if !tt.expectPanic && !tt.validate(got) {
				t.Errorf("Random.Generate() = %v, validation failed", got)
			}
		})
	}
}

func FuzzRandomGenerate(f *testing.F) {
	f.Fuzz(func(t *testing.T, in int) {
		t.Parallel()

		r := &acopw.Random{
			Length: in,
		}

		got := r.Generate()

		if in > 0 && len(got) != in {
			t.Errorf("Random.Generate() = %v, want %v", got, in)
		}
	})
}

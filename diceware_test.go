package acopw_test

import (
	"crypto/rand"
	"errors"
	"strings"
	"testing"

	"git.sr.ht/~jamesponddotco/acopw-go"
)

type errorReader struct{}

func (*errorReader) Read(_ []byte) (int, error) {
	return 0, errors.New("read error")
}

func TestDiceware_Generate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		validate func(string, error) bool
		name     string
		diceware *acopw.Diceware
	}{
		{
			name: "Default",
			diceware: &acopw.Diceware{
				Separator:  " ",
				Length:     0,
				Capitalize: true,
			},
			validate: func(generated string, err error) bool {
				if err != nil {
					return false
				}

				words := strings.Split(generated, " ")
				if len(words) != 8 {
					return false
				}

				capitalizedWordFound := false
				for _, word := range words {
					if word[0] >= 'A' && word[0] <= 'Z' {
						capitalizedWordFound = true
						break
					}
				}

				return capitalizedWordFound
			},
		},
		{
			name: "NoCapitalization",
			diceware: &acopw.Diceware{
				Rand:       rand.Reader,
				Separator:  " ",
				Length:     5,
				Capitalize: false,
			},
			validate: func(generated string, err error) bool {
				if err != nil {
					return false
				}

				words := strings.Split(generated, " ")
				if len(words) != 5 {
					return false
				}

				for _, word := range words {
					if word[0] >= 'A' && word[0] <= 'Z' {
						return false
					}
				}

				return true
			},
		},
		{
			name: "SeparatorError",
			diceware: &acopw.Diceware{
				Rand: &errorReader{},
			},
			validate: func(generated string, err error) bool {
				return err != nil
			},
		},
		{
			name: "CapitalizeError",
			diceware: &acopw.Diceware{
				Rand:       &errorReader{},
				Separator:  " ",
				Length:     5,
				Capitalize: true,
			},
			validate: func(generated string, err error) bool {
				return err != nil
			},
		},
		{
			name: "RandomElementError",
			diceware: &acopw.Diceware{
				Rand:      &errorReader{},
				Separator: " ",
			},
			validate: func(generated string, err error) bool {
				return err != nil
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.diceware.Generate()

			if !tt.validate(got, err) {
				t.Errorf("Diceware.Generate() = %v, error = %v", got, err)
			}
		})
	}
}

func FuzzDicewareRandom(f *testing.F) {
	f.Fuzz(func(t *testing.T, length int, capitalize bool) {
		diceware := &acopw.Diceware{
			Rand:       rand.Reader,
			Separator:  " ",
			Length:     length,
			Capitalize: capitalize,
		}

		got, err := diceware.Generate()
		if err != nil {
			t.Fatal(err)
		}

		if got == "" {
			t.Errorf("Diceware.Generate() = %v, want non-empty string", got)
		}

		if length > 1 && len(strings.Split(got, " ")) != length {
			t.Errorf("Diceware.Generate() = %v, want %d words", got, length)
		}
	})
}

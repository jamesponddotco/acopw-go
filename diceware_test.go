package acopw_test

import (
	"crypto/rand"
	"strings"
	"testing"

	"git.sr.ht/~jamesponddotco/acopw-go"
)

func TestDiceware_Generate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		validate func(string) bool
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
			validate: func(generated string) bool {
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
			validate: func(generated string) bool {
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
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := tt.diceware.Generate()
			if !tt.validate(got) {
				t.Errorf("Diceware.Generate() = %v, validation failed", got)
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

		got := diceware.Generate()
		if got == "" {
			t.Errorf("Diceware.Generate() = %v, want non-empty string", got)
		}

		if length > 1 && len(strings.Split(got, " ")) != length {
			t.Errorf("Diceware.Generate() = %v, want %d words", got, length)
		}
	})
}

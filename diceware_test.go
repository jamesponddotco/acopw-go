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
		name     string
		diceware acopw.Diceware
		validate func(string) bool
	}{
		{
			name: "Default",
			diceware: acopw.Diceware{
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
			diceware: acopw.Diceware{
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

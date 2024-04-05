package acopw_test

import (
	"strings"
	"testing"

	"git.sr.ht/~jamesponddotco/acopw-go"
)

func TestDiceware_Generate(t *testing.T) {
	t.Parallel()

	customWordList := []string{
		"custom",
		"word",
		"list",
	}

	tests := []struct {
		name     string
		diceware *acopw.Diceware
		validate func(string) bool
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
				if len(words) != acopw.DefaultDicewareLength {
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
		{
			name: "CustomSeparator",
			diceware: &acopw.Diceware{
				Separator: "-",
			},
			validate: func(generated string) bool {
				return strings.Contains(generated, "-")
			},
		},
		{
			name: "CustomWordList",
			diceware: &acopw.Diceware{
				Words:     customWordList,
				Length:    3,
				Separator: " ",
			},
			validate: func(generated string) bool {
				customWords := make(map[string]bool)
				for _, w := range customWordList {
					customWords[w] = true
				}

				for _, word := range strings.Split(generated, " ") {
					if !customWords[word] {
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
				t.Errorf("Diceware.Generate() = %v, want valid password", got)
			}
		})
	}
}

func FuzzDiceware_Generate(f *testing.F) {
	f.Fuzz(func(t *testing.T, length int, capitalize bool) {
		diceware := &acopw.Diceware{
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

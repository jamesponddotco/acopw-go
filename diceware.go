package acopw

import (
	"crypto/rand"
	_ "embed"
	"fmt"
	"io"
	"strings"

	"git.sr.ht/~jamesponddotco/acopw-go/internal/cryptoutil"
	"git.sr.ht/~jamesponddotco/xstd-go/xerrors"
	"git.sr.ht/~jamesponddotco/xstd-go/xstrings"
)

const (
	// ErrDicewarePassword is returned when a diceware password cannot be generated.
	ErrDicewarePassword xerrors.Error = "failed to generate diceware password"

	// ErrWordPool is returned when words cannot be retrieved from the pool.
	ErrWordPool xerrors.Error = "failed to get words from pool"
)

//go:embed words/word-list.txt
var _wordsData string

var (
	_words      = strings.Split(_wordsData, "\n")        //nolint:gochecknoglobals // we want this initialized with the package
	_separators = []string{"-", "_", ".", ";", " ", "/"} //nolint:gochecknoglobals // keeping it global avoids us redefining them every time
)

// DefaultDicewareLength is the default length of a diceware password.
const DefaultDicewareLength int = 8

// Diceware contains configuration options for generating a diceware password.
type Diceware struct {
	// Rand provides the source of entropy for generating the diceware
	// password. If Rand is nil, the cryptographic random reader in package
	// crypto/rand is used.
	Rand io.Reader

	// Separator is the string used to separate words in the password.
	Separator string

	// Length is the number of words in the password.
	Length int

	// Capitalize indicates whether a random word should be capitalized.
	Capitalize bool
}

// Generate generates a diceware password.
func (d *Diceware) Generate() (string, error) {
	if d.Length < 1 {
		d.Length = DefaultDicewareLength
	}

	if d.Separator == "" {
		separator, err := cryptoutil.RandomElement(d.reader(), _separators)
		if err != nil {
			return "", fmt.Errorf("%w: %w", ErrDicewarePassword, err)
		}

		d.Separator = separator
	}

	words := make([]string, 0, d.Length)

	for i := 0; i < d.Length; i++ {
		element, err := cryptoutil.RandomElement(d.reader(), _words)
		if err != nil {
			return "", fmt.Errorf("%w: %w", ErrDicewarePassword, err)
		}

		words = append(words, element)
	}

	if d.Capitalize {
		capitalizeIndex, err := cryptoutil.Int(d.reader(), len(words))
		if err != nil {
			return "", fmt.Errorf("%w: %w", ErrDicewarePassword, err)
		}

		words[capitalizeIndex] = strings.ToUpper(words[capitalizeIndex])
	}

	return xstrings.JoinWithSeparator(d.Separator, words...), nil
}

// reader returns the reader to use for generating the diceware password.
func (d *Diceware) reader() io.Reader {
	if d.Rand != nil {
		return d.Rand
	}

	return rand.Reader
}

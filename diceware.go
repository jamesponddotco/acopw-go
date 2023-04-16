package acopw

import (
	"crypto/rand"
	_ "embed"
	"io"
	"strings"
	"sync"

	"git.sr.ht/~jamesponddotco/xstd-go/xcrypto/xrand"
	"git.sr.ht/~jamesponddotco/xstd-go/xstrings"
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
	// wordsPool is the pool of words to use for generating the password.
	wordsPool sync.Pool

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
func (d *Diceware) Generate() string {
	if d.Length < 1 {
		d.Length = DefaultDicewareLength
	}

	if d.Separator == "" {
		d.Separator = d.randomElement(_separators)
	}

	if d.wordsPool.New == nil {
		d.wordsPool.New = func() any {
			return make([]string, 0, DefaultDicewareLength)
		}
	}

	words, ok := d.wordsPool.Get().([]string)
	if !ok {
		return ""
	}

	words = words[:d.Length]
	defer d.wordsPool.Put(&words)

	for i := 0; i < d.Length; i++ {
		words[i] = d.randomElement(_words)
	}

	xrand.Shuffle(words, d.reader())

	// Capitalize a random word if required.
	if d.Capitalize {
		capitalizeIndex := xrand.IntChaChaCha(len(words), d.reader())
		words[capitalizeIndex] = strings.ToUpper(words[capitalizeIndex])
	}

	return xstrings.JoinWithSeparator(d.Separator, words...)
}

// reader returns the reader to use for generating the diceware password.
func (d *Diceware) reader() io.Reader {
	if d.Rand != nil {
		return d.Rand
	}

	return rand.Reader
}

// randomElement returns a random element from the given string silce.
func (d *Diceware) randomElement(elements []string) string {
	index := xrand.IntChaChaCha(len(elements), d.reader())

	return elements[index]
}

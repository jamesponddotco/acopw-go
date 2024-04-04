package acopw

import (
	"crypto/rand"
	_ "embed"
	"io"
	mrand "math/rand/v2"
	"strings"

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
	// random provides the source of entropy for generating the diceware password.
	random *mrand.Rand

	// Separator is the string used to separate words in the password.
	Separator string

	// Words is the list of words used to generate the password. If the list is
	// empty, the default word list is used.
	Words []string

	// Length is the number of words in the password.
	Length int

	// Capitalize indicates whether a random word should be capitalized.
	Capitalize bool
}

// Generate generates a diceware password.
func (d *Diceware) Generate() string {
	if d.random == nil {
		var seed [32]byte

		if _, err := io.ReadFull(rand.Reader, seed[:]); err != nil {
			panic(err)
		}

		d.random = mrand.New(mrand.NewChaCha8(seed))
	}

	if d.Length < 1 {
		d.Length = DefaultDicewareLength
	}

	var (
		index    int
		wordList = d.Words
	)

	if len(wordList) == 0 {
		wordList = _words
	}

	if d.Separator == "" {
		index = d.random.IntN(len(_separators))

		d.Separator = _separators[index]
	}

	capitalizeIndex := -1

	if d.Capitalize {
		capitalizeIndex = d.random.IntN(d.Length)
	}

	words := make([]string, 0, d.Length)

	for i := 0; i < d.Length; i++ {
		var (
			index = d.random.IntN(len(wordList))
			word  = wordList[index]
		)

		if i == capitalizeIndex {
			word = strings.ToUpper(word)
		}

		words = append(words, word)
	}

	return xstrings.JoinWithSeparator(d.Separator, words...)
}

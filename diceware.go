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

// Diceware is a policy for generating ChaCha8-based cryptographically strong
// diceware passwords.
type Diceware struct {
	// random provides the source of entropy for generating the diceware password.
	random *mrand.Rand

	// Separator is the string used to separate words in the password.
	Separator string

	// Words is a list of words included in the generated password. If the list
	// is empty, a default word list is used.
	Words []string

	// Length is the number of words to include in the generated password. If
	// less than 1, it defaults to 8.
	Length int

	// Capitalize indicates whether a random word should be capitalized.
	Capitalize bool
}

// Generate returns a cryptographically strong diceware password for the policy.
// It panics if it can't get entropy from the source of randomness.
func (d *Diceware) Generate() string {
	if d.random == nil {
		var seed [32]byte

		if _, err := io.ReadFull(rand.Reader, seed[:]); err != nil {
			panic(err)
		}

		d.random = mrand.New(mrand.NewChaCha8(seed)) //nolint:gosec // we seed with crypto/rand
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
		index = d.random.IntN(len(wordList))
		word := wordList[index]

		if i == capitalizeIndex {
			word = strings.ToUpper(word)
		}

		words = append(words, word)
	}

	return xstrings.JoinWithSeparator(d.Separator, words...)
}

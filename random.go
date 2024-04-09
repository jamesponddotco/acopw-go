package acopw

import (
	"crypto/rand"
	"io"
	mrand "math/rand/v2"

	"git.sr.ht/~jamesponddotco/xstd-go/xerrors"
	"git.sr.ht/~jamesponddotco/xstd-go/xstrings"
)

// ErrInvalidCharset is returned when the charset is invalid.
const ErrInvalidCharset xerrors.Error = "no characters to build password in the charset"

// DefaultRandomLength is the default length of a random password.
const DefaultRandomLength int = 128

// Random contains configuration options for generating a random password.
type Random struct {
	// random provides the source of entropy for generating the password.
	random *mrand.Rand

	// characters is the character set to use for generating the password.
	characters []string

	// ExcludedCharset is a list of characters that should not be used in the password.
	ExcludedCharset []string

	// Length is the length of the password.
	Length int

	// UseLower, UseUpper, UseNumbers, and UseSymbols specify whether or not to use the corresponding character class.
	UseLower   bool
	UseUpper   bool
	UseNumbers bool
	UseSymbols bool
}

// Generate generates a random password.
func (r *Random) Generate() string { //nolint:unparam // appears to be a false positive
	if r.random == nil {
		var seed [32]byte

		if _, err := io.ReadFull(rand.Reader, seed[:]); err != nil {
			panic(err)
		}

		r.random = mrand.New(mrand.NewChaCha8(seed))
	}

	if r.Length < 1 {
		r.Length = DefaultRandomLength
	}

	if !r.UseLower && !r.UseUpper && !r.UseNumbers && !r.UseSymbols {
		r.UseLower = true
		r.UseUpper = true
		r.UseNumbers = true
		r.UseSymbols = true
	}

	charset := r.charset()

	if len(charset) == 0 {
		panic(ErrInvalidCharset)
	}

	password := make([]string, 0, r.Length)

	for i := 0; i < r.Length; i++ {
		var (
			index = r.random.IntN(len(charset))
			char  = charset[index]
		)

		password = append(password, char)
	}

	return xstrings.Join(password...)
}

// Charset returns the character set to use for generating the password.
func (r *Random) charset() []string {
	if r.characters == nil { //nolint:nestif // what other way is there?
		charset := make([]string, 0, len(_charsetLower)+len(_charsetUpper)+len(_charsetNumbers)+len(_charsetSymbols))

		if r.UseLower {
			charset = append(charset, _charsetLower...)
		}

		if r.UseUpper {
			charset = append(charset, _charsetUpper...)
		}

		if r.UseNumbers {
			charset = append(charset, _charsetNumbers...)
		}

		if r.UseSymbols {
			charset = append(charset, _charsetSymbols...)
		}

		if len(r.ExcludedCharset) > 0 {
			excludedChars := make(map[string]bool, len(r.ExcludedCharset))
			for _, char := range r.ExcludedCharset {
				excludedChars[char] = true
			}

			filteredCharset := make([]string, 0, len(charset))
			for _, char := range charset { //nolint:wsl // looks like a false positive to me
				if !excludedChars[char] {
					filteredCharset = append(filteredCharset, char)
				}
			}

			charset = filteredCharset
		}

		r.characters = charset
	}

	return r.characters
}

package acopw

import (
	"crypto/rand"
	"io"
	mrand "math/rand/v2"
	"strings"
)

// DefaultRandomLength is the default length of a random password.
const DefaultRandomLength int = 128

// Random is a policy for generating ChaCha8-based cryptographically strong
// random passwords. Instances are not safe for concurrent use.
type Random struct {
	// random provides the source of entropy for generating the password.
	random *mrand.Rand

	// characters is the character set to use for generating the password.
	characters []string

	// ExcludedCharset is a list of characters that should not be included in
	// the generated password.
	ExcludedCharset []string

	// Length is the length of the password. If less than 1, it defaults to 128.
	//
	// It's the caller's responsibility to limit the maximum length to prevent
	// memory exhaustion via large length values.
	Length int

	// UseLower, UseUpper, UseNumbers, and UseSymbols specify whether or not to
	// use the corresponding character class in the generated password.
	//
	// If none of these are true, it defaults to true for all four.
	UseLower   bool
	UseUpper   bool
	UseNumbers bool
	UseSymbols bool
}

// Generate returns a cryptographically strong random password for the policy.
// It panics if it can't get entropy from the source of randomness or if the
// internally generated character set is empty.
func (r *Random) Generate() string { //nolint:unparam // appears to be a false positive
	if r.random == nil {
		var seed [32]byte

		if _, err := io.ReadFull(rand.Reader, seed[:]); err != nil {
			panic(err)
		}

		r.random = mrand.New(mrand.NewChaCha8(seed)) //nolint:gosec // we seed with crypto/rand
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
		panic("no characters to build password in the charset")
	}

	password := make([]string, 0, r.Length)

	for range r.Length {
		var (
			index = r.random.IntN(len(charset))
			char  = charset[index]
		)

		password = append(password, char)
	}

	return strings.Join(password, "")
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

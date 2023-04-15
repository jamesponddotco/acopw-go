package acopw

import (
	"crypto/rand"
	"io"
	"sync"

	"git.sr.ht/~jamesponddotco/xstd-go/xcrypto/xrand"
	"git.sr.ht/~jamesponddotco/xstd-go/xerrors"
	"git.sr.ht/~jamesponddotco/xstd-go/xstrings"
)

const (

	// ErrNoCharacterClasses is returned when no character classes are used.
	ErrNoCharacterClasses xerrors.Error = "at least one character class must be used"
)

const (
	Lowercase = xstrings.LowercaseLetters
	Uppercase = xstrings.UppercaseLetters
	Numbers   = xstrings.Numbers
	Symbols   = xstrings.Symbols
)

// DefaultRandomLength is the default length of a random password.
const DefaultRandomLength int = 128

// Random contains configuration options for generating a random password.
type Random struct {
	// Charset is the character set to use for generating the password.
	charset string

	// Rand provides the source of entropy for generating the password. If Rand
	// is nil, the cryptographic random reader in package crypto/rand is used.
	Rand io.Reader

	// ExcludedCharset is a list of characters that should not be used in the password.
	ExcludedCharset []string

	// Length is the length of the password.
	Length int

	// UseLower, UseUpper, UseNumbers, and UseSymbols specify whether or not to use the corresponding character class.
	UseLower   bool
	UseUpper   bool
	UseNumbers bool
	UseSymbols bool

	// once is used to ensure that the charset is only generated once.
	once sync.Once
}

// Generate generates a random password.
func (r *Random) Generate() string {
	if r.Length < 1 {
		r.Length = DefaultRandomLength
	}

	if !r.UseLower && !r.UseUpper && !r.UseNumbers && !r.UseSymbols {
		r.UseLower = true
		r.UseUpper = true
		r.UseNumbers = true
		r.UseSymbols = true
	}

	var (
		charset    = r.Charset()
		reader     = r.reader()
		password   = make([]byte, r.Length)
		bufferSize = int(float64(r.Length) * 1.3)
	)

	for i, j, randomBytes := 0, 0, []byte{}; i < r.Length; j++ {
		if j%bufferSize == 0 {
			randomBytes = xrand.BytesWithReader(bufferSize, reader)
		}

		if idx := int(randomBytes[j%bufferSize] & _indexMask); idx < len(charset) {
			password[i] = charset[idx]
			i++
		}
	}

	return string(password)
}

// Charset returns the character set to use for generating the password.
func (r *Random) Charset() string {
	r.once.Do(func() {
		var charset string

		if r.UseLower {
			charset += Lowercase
		}

		if r.UseUpper {
			charset += Uppercase
		}

		if r.UseNumbers {
			charset += Numbers
		}

		if r.UseSymbols {
			charset += Symbols
		}

		for _, excluded := range r.ExcludedCharset {
			charset = xstrings.Remove(charset, excluded)
		}

		r.charset = charset
	})

	return r.charset
}

// reader returns the source of entropy for generating the password.
func (r *Random) reader() io.Reader {
	if r.Rand != nil {
		return r.Rand
	}

	return rand.Reader
}

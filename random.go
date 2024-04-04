package acopw

import (
	"crypto/rand"
	"io"
	"sync"

	"git.sr.ht/~jamesponddotco/xstd-go/xerrors"
	"git.sr.ht/~jamesponddotco/xstd-go/xstrings"
	"git.sr.ht/~jamesponddotco/xstd-go/xunsafe"
)

// ErrInvalidCharset is returned when the charset is invalid.
const ErrInvalidCharset xerrors.Error = "no characters to build password in the charset"

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

	charset := r.Charset()
	if charset == "" {
		panic(ErrInvalidCharset)
	}

	var (
		reader      = rand.Reader
		password    = make([]byte, r.Length)
		randomBytes = make([]byte, r.Length)
		maxByte     = byte(256 - (256 % len(charset)))
	)

	if _, err := io.ReadFull(reader, randomBytes); err != nil {
		panic(err)
	}

	for i := 0; i < r.Length; i++ {
		b := randomBytes[i]

		for b >= maxByte {
			if _, err := io.ReadFull(reader, randomBytes[i:i+1]); err != nil {
				panic(err)
			}

			b = randomBytes[i]
		}

		password[i] = charset[int(b)%len(charset)]
	}

	return xunsafe.BytesToString(password)
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

		if len(r.ExcludedCharset) > 0 {
			for _, excluded := range r.ExcludedCharset {
				charset = xstrings.Remove(charset, excluded)
			}
		}

		r.charset = charset
	})

	return r.charset
}

package acopw

import (
	"crypto/rand"
	"io"
	"strings"

	"git.sr.ht/~jamesponddotco/xstd-go/xcrypto/xrand"
	"git.sr.ht/~jamesponddotco/xstd-go/xstrings"
	"git.sr.ht/~jamesponddotco/xstd-go/xunsafe"
)

const (
	_RandomIndexBits = 7
	_RandomIndexMask = byte(1<<_RandomIndexBits - 1)
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
		charset     = r.Charset()
		charsetLen  = len(charset)
		reader      = r.reader()
		password    = make([]byte, r.Length)
		randomBytes = xrand.BytesWithReader(r.Length, reader)
	)

	for i, j := 0, 0; i < r.Length; j++ {
		if idx := int(randomBytes[j%r.Length] & _RandomIndexMask); idx < charsetLen {
			password[i] = charset[idx]
			i++
		}
	}

	return xunsafe.BytesToString(password)
}

// Charset returns the character set to use for generating the password.
func (r *Random) Charset() string {
	if r.charset == "" {
		var builder strings.Builder

		builder.Grow(len(Lowercase) + len(Uppercase) + len(Numbers) + len(Symbols))

		if r.UseLower {
			builder.WriteString(Lowercase)
		}

		if r.UseUpper {
			builder.WriteString(Uppercase)
		}

		if r.UseNumbers {
			builder.WriteString(Numbers)
		}

		if r.UseSymbols {
			builder.WriteString(Symbols)
		}

		if len(r.ExcludedCharset) > 0 {
			for _, excluded := range r.ExcludedCharset {
				r.charset = xstrings.Remove(builder.String(), excluded)
			}
		} else {
			r.charset = builder.String()
		}
	}

	return r.charset
}

// reader returns the source of entropy for generating the password.
func (r *Random) reader() io.Reader {
	if r.Rand != nil {
		return r.Rand
	}

	return rand.Reader
}

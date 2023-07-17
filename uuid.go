package acopw

import (
	"crypto/rand"
	"fmt"
	"io"

	"git.sr.ht/~jamesponddotco/xstd-go/xunsafe"
)

const (
	// _hexDigits is a lookup table for converting a byte into a hex string.
	_hexDigits string = "0123456789abcdef"

	// _separator is the separator used in the UUID string.
	_separator byte = '-'
)

// UUID contains configuration options for generating a 128-bit Universal Unique
// Identifier (UUID) as defined in [RFC 4122]. Specifically, it generates the
// random variant, UUIDv4.
//
// [RFC 4122]: https://tools.ietf.org/html/rfc4122
type UUID struct {
	// Rand provides the source of entropy for generating the UUID. If Rand
	// is nil, the cryptographic random reader in package crypto/rand is used.
	Rand io.Reader

	// bytes is a byte slice of length 16 that is used as a buffer for the UUID.
	bytes [16]byte
}

// Generate generates a random UUIDv4.
func (u *UUID) Generate() (string, error) {
	if _, err := io.ReadFull(u.reader(), u.bytes[:]); err != nil {
		return "", fmt.Errorf("%w", err)
	}

	u.bytes[6] = (u.bytes[6] & 0x0f) | 0x40
	u.bytes[8] = (u.bytes[8] & 0x3f) | 0x80

	buf := make([]byte, 0, 36)

	for i := 0; i < 16; i++ {
		if i == 4 || i == 6 || i == 8 || i == 10 {
			buf = append(buf, _separator)
		}

		buf = append(buf, _hexDigits[u.bytes[i]>>4], _hexDigits[u.bytes[i]&0x0f])
	}

	return xunsafe.BytesToString(buf), nil
}

// reader returns the source of entropy for generating the UUID.
func (u *UUID) reader() io.Reader {
	if u.Rand != nil {
		return u.Rand
	}

	return rand.Reader
}

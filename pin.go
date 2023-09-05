package acopw

import (
	"crypto/rand"
	"io"

	"git.sr.ht/~jamesponddotco/xstd-go/xstrings"
	"git.sr.ht/~jamesponddotco/xstd-go/xunsafe"
)

// DefaultPINLength is the default length of a PIN.
const DefaultPINLength int = 6

// PIN contains configuration options for generating PIN pins.
type PIN struct {
	// Rand provides the source of entropy for generating the PIN. If Rand is
	// nil, the cryptographic random reader in package crypto/rand is used.
	Rand io.Reader

	// Length is the length of the generated PIN.
	Length int
}

// Generate generates a random PIN.
func (p *PIN) Generate() string {
	if p.Length < 1 {
		p.Length = DefaultPINLength
	}

	var (
		charset     = xstrings.Numbers
		reader      = p.reader()
		pin         = make([]byte, p.Length)
		randomBytes = make([]byte, p.Length)
		maxByte     = byte(256 - (256 % len(charset)))
	)

	_, err := io.ReadFull(reader, randomBytes)
	if err != nil {
		return ""
	}

	for i := 0; i < p.Length; i++ {
		b := randomBytes[i]
		if b >= maxByte {
			continue
		}

		pin[i] = charset[int(b)%len(charset)]
	}

	return xunsafe.BytesToString(pin)
}

// reader returns the reader to use for generating the PIN.
func (p *PIN) reader() io.Reader {
	if p.Rand != nil {
		return p.Rand
	}

	return rand.Reader
}

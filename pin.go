package acopw

import (
	"crypto/rand"
	"io"
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

	pin := &Random{
		Length:     p.Length,
		Rand:       p.reader(),
		UseNumbers: true,
	}

	return pin.Generate()
}

// reader returns the reader to use for generating the PIN.
func (p *PIN) reader() io.Reader {
	if p.Rand != nil {
		return p.Rand
	}

	return rand.Reader
}

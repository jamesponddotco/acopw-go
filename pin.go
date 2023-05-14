package acopw

import (
	"crypto/rand"
	"io"

	"git.sr.ht/~jamesponddotco/xstd-go/xcrypto/xrand"
	"git.sr.ht/~jamesponddotco/xstd-go/xstrings"
	"git.sr.ht/~jamesponddotco/xstd-go/xunsafe"
)

const (
	_PINIndexBits = 4
	_PINIndexMask = 1<<_PINIndexBits - 1
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
		charset    = xstrings.Numbers
		reader     = p.reader()
		pin        = make([]byte, p.Length)
		bufferSize = int(float64(p.Length) * 1.3)
	)

	for i, j, randomBytes := 0, 0, []byte{}; i < p.Length; j++ {
		if j%bufferSize == 0 {
			randomBytes = xrand.BytesWithReader(bufferSize, reader)
		}

		if idx := int(randomBytes[j%bufferSize] & _PINIndexMask); idx < len(charset) {
			pin[i] = charset[idx]
			i++
		}
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

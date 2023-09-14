package acopw

import (
	"crypto/rand"
	"fmt"
	"io"

	"git.sr.ht/~jamesponddotco/xstd-go/xerrors"
	"git.sr.ht/~jamesponddotco/xstd-go/xstrings"
	"git.sr.ht/~jamesponddotco/xstd-go/xunsafe"
)

// ErrRandomPIN is returned when a random PIN cannot be generated.
const ErrRandomPIN xerrors.Error = "unable to generate random PIN"

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
func (p *PIN) Generate() (string, error) {
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
		return "", fmt.Errorf("%w: %w", ErrRandomPIN, err)
	}

	for i := 0; i < p.Length; i++ {
		b := randomBytes[i]
		if b >= maxByte {
			_, err := io.ReadFull(reader, randomBytes[i:i+1])
			if err != nil {
				return "", fmt.Errorf("%w: %w", ErrRandomPIN, err)
			}

			b = randomBytes[i]
		}

		pin[i] = charset[int(b)%len(charset)]
	}

	return xunsafe.BytesToString(pin), nil
}

// reader returns the reader to use for generating the PIN.
func (p *PIN) reader() io.Reader {
	if p.Rand != nil {
		return p.Rand
	}

	return rand.Reader
}

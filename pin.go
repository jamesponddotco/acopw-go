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
		reader      = rand.Reader
		pin         = make([]byte, p.Length)
		randomBytes = make([]byte, p.Length)
		maxByte     = byte(256 - (256 % len(charset)))
	)

	if _, err := io.ReadFull(reader, randomBytes); err != nil {
		panic(err)
	}

	for i := 0; i < p.Length; i++ {
		b := randomBytes[i]
		if b >= maxByte {
			if _, err := io.ReadFull(reader, randomBytes[i:i+1]); err != nil {
				panic(err)
			}

			b = randomBytes[i]
		}

		pin[i] = charset[int(b)%len(charset)]
	}

	return xunsafe.BytesToString(pin)
}

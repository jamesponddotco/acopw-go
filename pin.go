package acopw

import (
	"crypto/rand"
	"io"
	mrand "math/rand/v2"

	"git.sr.ht/~jamesponddotco/xstd-go/xstrings"
)

// DefaultPINLength is the default length of a PIN.
const DefaultPINLength int = 6

// PIN contains configuration options for generating PIN pins.
type PIN struct {
	// random provides the source of entropy for generating the PIN.
	random *mrand.Rand

	// Length is the length of the generated PIN.
	Length int
}

// Generate generates a random PIN.
func (p *PIN) Generate() string {
	if p.random == nil {
		var seed [32]byte

		if _, err := io.ReadFull(rand.Reader, seed[:]); err != nil {
			panic(err)
		}

		p.random = mrand.New(mrand.NewChaCha8(seed))
	}

	if p.Length < 1 {
		p.Length = DefaultPINLength
	}

	pin := make([]string, 0, p.Length)

	for i := 0; i < p.Length; i++ {
		var (
			index = p.random.IntN(len(_charsetNumbers))
			char  = _charsetNumbers[index]
		)

		pin = append(pin, char)
	}

	return xstrings.Join(pin...)
}

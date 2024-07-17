package acopw

import (
	"crypto/rand"
	"io"
	mrand "math/rand/v2"

	"git.sr.ht/~jamesponddotco/xstd-go/xstrings"
)

// DefaultPINLength is the default length of a PIN.
const DefaultPINLength int = 6

// PIN is a policy for generating ChaCha8-based cryptographically strong random
// PINs.
type PIN struct {
	// random provides the source of entropy for generating the PIN.
	random *mrand.Rand

	// Length is the length of the generated PIN. If less than 1, it defaults to
	// 6.
	Length int
}

// Generate returns a cryptographically strong PIN for the policy. It panics if
// it can't get entropy from the source of randomness.
func (p *PIN) Generate() string {
	if p.random == nil {
		var seed [32]byte

		if _, err := io.ReadFull(rand.Reader, seed[:]); err != nil {
			panic(err)
		}

		p.random = mrand.New(mrand.NewChaCha8(seed)) //nolint:gosec // we seed with crypto/rand
	}

	if p.Length < 1 {
		p.Length = DefaultPINLength
	}

	pin := make([]string, 0, p.Length)

	for range p.Length {
		var (
			index = p.random.IntN(len(_charsetNumbers))
			char  = _charsetNumbers[index]
		)

		pin = append(pin, char)
	}

	return xstrings.Join(pin...)
}

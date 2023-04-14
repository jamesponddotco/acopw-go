package acopw

import (
	"crypto/rand"
	"encoding/binary"
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

	var (
		reader = p.reader()
		pin    = make([]byte, p.Length)
	)

	for iter, cache, remain := p.Length-1, int64(0), 0; iter >= 0; {
		if remain == 0 {
			err := binary.Read(reader, binary.BigEndian, &cache)
			if err != nil {
				panic(err)
			}

			remain = _indexMax
		}

		if index := int(cache & _indexMask); index < len(Numbers) {
			pin[iter] = Numbers[index]
			iter--
		}

		cache >>= _indexBits
		remain--
	}

	return string(pin)
}

// reader returns the reader to use for generating the PIN.
func (p *PIN) reader() io.Reader {
	if p.Rand != nil {
		return p.Rand
	}

	return rand.Reader
}

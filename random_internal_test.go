package acopw

import (
	"crypto/rand"
	"testing"
)

func TestRandom_Generate_Panic_ReadFull(t *testing.T) { //nolint:paralleltest // we're modifying rand.Reader
	origReader := rand.Reader

	t.Cleanup(func() {
		rand.Reader = origReader
	})

	rand.Reader = &mockReader{}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected Generate to panic due to rand.Reader error, but it did not")
		}
	}()

	random := Random{}

	_ = random.Generate()
}

func TestRandom_Generate_Panic_EmptyCharset(t *testing.T) {
	t.Parallel()

	charset := make([]string, 0, len(_charsetLower)+len(_charsetUpper)+len(_charsetNumbers)+len(_charsetSymbols))
	charset = append(charset, _charsetLower...)
	charset = append(charset, _charsetUpper...)
	charset = append(charset, _charsetNumbers...)
	charset = append(charset, _charsetSymbols...)

	random := Random{
		ExcludedCharset: charset,
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected Generate to panic due to empty charset, but it did not")
		}
	}()

	_ = random.Generate()
}

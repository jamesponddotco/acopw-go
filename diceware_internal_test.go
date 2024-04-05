package acopw

import (
	"crypto/rand"
	"errors"
	"testing"
)

type mockReader struct{}

func (*mockReader) Read(_ []byte) (n int, err error) {
	return 0, errors.New("mock read error")
}

func TestDiceware_Generate_Panic(t *testing.T) { //nolint:paralleltest // we're modifying rand.Reader
	origReader := rand.Reader

	t.Cleanup(func() {
		rand.Reader = origReader
	})

	rand.Reader = &mockReader{}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected Generate to panic, but it did not")
		}
	}()

	diceware := Diceware{
		Length:     7,    // Use 7 words.
		Capitalize: true, // Capitalize the first letter of a random word.
	}

	_ = diceware.Generate()
}

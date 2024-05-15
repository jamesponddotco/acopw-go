package acopw

import (
	"crypto/rand"
	"testing"

	"git.sr.ht/~jamesponddotco/xstd-go/xerrors"
)

const ErrGeneric xerrors.Error = "error"

type mockReader struct{}

func (*mockReader) Read(_ []byte) (n int, err error) {
	return 0, ErrGeneric
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
		Length:     7,
		Capitalize: true,
	}

	_ = diceware.Generate()
}

package acopw

import (
	"crypto/rand"
	"testing"
)

func TestPIN_Generate_Panic(t *testing.T) { //nolint:paralleltest // we're modifying rand.Reader
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

	pin := PIN{}

	_ = pin.Generate()
}

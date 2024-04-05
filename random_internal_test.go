package acopw

import (
	"crypto/rand"
	"testing"
)

func TestRandom_Generate_Panic_FirstReadFull(t *testing.T) { //nolint:paralleltest // we're modifying rand.Reader
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

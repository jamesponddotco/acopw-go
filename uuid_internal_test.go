package acopw

import (
	"crypto/rand"
	"testing"
)

func TestUUID_Generate_Panic(t *testing.T) { //nolint:paralleltest // we're modifying rand.Reader
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

	uuid := &UUID{}
	_ = uuid.Generate()
}

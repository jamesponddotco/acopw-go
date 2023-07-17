package acopw_test

import (
	"testing"

	"git.sr.ht/~jamesponddotco/acopw-go"
)

func BenchmarkUUID_Generate(b *testing.B) {
	var uuid acopw.UUID

	for i := 0; i < b.N; i++ {
		_, err := uuid.Generate()
		if err != nil {
			b.Fatal(err)
		}
	}
}

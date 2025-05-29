package acopw_test

import (
	"testing"

	"git.sr.ht/~jamesponddotco/acopw-go"
)

func BenchmarkPIN_Generate(b *testing.B) {
	pin := acopw.PIN{
		Length: 6,
	}

	b.ResetTimer()

	for range b.N {
		_ = pin.Generate()
	}
}

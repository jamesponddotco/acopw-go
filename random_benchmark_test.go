package acopw_test

import (
	"testing"

	"git.sr.ht/~jamesponddotco/acopw-go"
)

func BenchmarkRandom_Generate(b *testing.B) {
	password := acopw.Random{
		ExcludedCharset: []string{
			" ",
			"&",
		},
		Length:     64,
		UseLower:   true,
		UseUpper:   true,
		UseSymbols: true,
	}

	b.ResetTimer()

	for range b.N {
		_ = password.Generate()
	}
}

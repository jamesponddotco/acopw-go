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

	for i := 0; i < b.N; i++ {
		_, err := password.Generate()
		if err != nil {
			b.Fatal(err)
		}
	}
}

package acopw_test

import (
	"testing"

	"git.sr.ht/~jamesponddotco/acopw-go"
)

func BenchmarkDiceware_Generate(b *testing.B) {
	diceware := &acopw.Diceware{
		Length:     7,
		Separator:  " ",
		Capitalize: true,
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = diceware.Generate()
	}
}

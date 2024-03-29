package acopw_test

import (
	"log"

	"git.sr.ht/~jamesponddotco/acopw-go"
)

func ExampleRandom_Generate() {
	// Define password policy.
	password := acopw.Random{
		ExcludedCharset: []string{
			" ", // Exclude spaces
			"&", // Exclude ampersands
		},
		Length:     64,   // Generate a 64 character password
		UseLower:   true, // Use lowercase letters
		UseUpper:   true, // Use uppercase letters
		UseSymbols: true, // Use symbols
	}

	// Generate a random password.
	generated, err := password.Generate()
	if err != nil {
		log.Fatal(err)
	}

	log.Print(generated)
	// Output:
	//
}

package acopw_test

import (
	"fmt"
	"log"
	"os"

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

	// Generate and print a random password. Use os.Stdout in the real world.
	if _, err := fmt.Fprintln(os.Stderr, password.Generate()); err != nil {
		log.Fatal(err)
	}
	// Output:
	//
}

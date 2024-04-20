package acopw_test

import (
	"fmt"
	"log"
	"os"

	"git.sr.ht/~jamesponddotco/acopw-go"
)

func ExampleDiceware_Generate() {
	// Define password policy.
	password := acopw.Diceware{
		Length:     7,    // Use 7 words.
		Capitalize: true, // Capitalize the first letter of a random word.
	}

	// Generate and print a random diceware password. Use os.Stdout in the real
	// world.
	if _, err := fmt.Fprintln(os.Stderr, password.Generate()); err != nil {
		log.Fatal(err)
	}
	// Output:
}

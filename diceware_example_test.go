package acopw_test

import (
	"log"

	"git.sr.ht/~jamesponddotco/acopw-go"
)

func ExampleDiceware_Generate() {
	// Define password policy.
	password := acopw.Diceware{
		Length:     7,    // Use 7 words.
		Capitalize: true, // Capitalize the first letter of a random word.
	}

	// Generate a password.
	pass, err := password.Generate()
	if err != nil {
		log.Fatal(err)
	}

	// Print the password.
	log.Print(pass)
	// Output:
}

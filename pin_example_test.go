package acopw_test

import (
	"log"

	"git.sr.ht/~jamesponddotco/acopw-go"
)

func ExamplePIN_Generate() {
	// Define your PIN policy.
	pin := acopw.PIN{
		Length: 6, // Generate a 6 digit PIN.
	}

	// Generate a random PIN.
	log.Print(pin.Generate())
	// Output:
	//
}

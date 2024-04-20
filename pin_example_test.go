package acopw_test

import (
	"fmt"
	"log"
	"os"

	"git.sr.ht/~jamesponddotco/acopw-go"
)

func ExamplePIN_Generate() {
	// Define your PIN policy.
	pin := acopw.PIN{
		Length: 6, // Generate a 6 digit PIN.
	}

	// Generate and print a random PIN. Use os.Stdout in the real world.
	if _, err := fmt.Fprintln(os.Stderr, pin.Generate()); err != nil {
		log.Fatal(err)
	}
	// Output:
	//
}

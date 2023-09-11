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

	// Generate a PIN.
	generated, err := pin.Generate()
	if err != nil {
		log.Fatal(err)
	}

	log.Print(generated)
	// Output:
	//
}

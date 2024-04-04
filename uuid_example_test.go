package acopw_test

import (
	"log"

	"git.sr.ht/~jamesponddotco/acopw-go"
)

func ExampleUUID_Generate() {
	var uuid acopw.UUID

	// Generate and print a random UUIDv4.
	log.Print(uuid.Generate())
	// Output:
	//
}

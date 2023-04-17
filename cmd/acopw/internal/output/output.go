// Package output provides a way to output a clean password or error.
package output

import (
	"log"
	"os"

	"golang.org/x/term"
)

// Password outputs a clean password.
func Password(password string) {
	plainLogger := log.New(os.Stdout, "", 0)

	switch {
	case term.IsTerminal(int(os.Stdout.Fd())):
		plainLogger.Print(password)
	default:
		plainLogger.Println(password)
	}
}

// Error outputs an error to stderr.
func Error(v ...any) {
	errorLogger := log.New(os.Stderr, "", 0)

	switch {
	case term.IsTerminal(int(os.Stderr.Fd())):
		errorLogger.Print(v...)
	default:
		errorLogger.Println(v...)
	}
}

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
	PasswordWithLogger(password, plainLogger)
}

// PasswordWithLogger outputs a clean password with a custom logger.
func PasswordWithLogger(password string, logger *log.Logger) {
	switch {
	case term.IsTerminal(int(os.Stdout.Fd())):
		logger.Print(password)
	default:
		logger.Println(password)
	}
}

// Error outputs an error to stderr.
func Error(v ...any) {
	errorLogger := log.New(os.Stderr, "", 0)
	ErrorWithLogger(v, errorLogger)
}

// ErrorWithLogger outputs an error with a custom logger.
func ErrorWithLogger(v any, logger *log.Logger) {
	switch {
	case term.IsTerminal(int(os.Stderr.Fd())):
		logger.Print(v)
	default:
		logger.Println(v)
	}
}

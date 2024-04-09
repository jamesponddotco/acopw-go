// Package acopw provides a simple way to generate cryptographically secure
// random and diceware passwords, and PINs.
package acopw

//nolint:gochecknoglobals // we want these initialized with the package
var (
	_charsetLower = []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	}
	_charsetUpper = []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	}
	_charsetNumbers = []string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	}
	_charsetSymbols = []string{
		"!", "\"", "#", "$", "%", "&", "'", "(", ")", "*", "+", ",", "-", ".", "/", ":", ";", "<", "=", ">", "?", "@", "[", "\\", "]", "^", "_", "{", "|", "}", "~",
	}
)

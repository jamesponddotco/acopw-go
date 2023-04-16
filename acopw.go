// Package acopw provides a simple way to generate cryptographically secure
// random and diceware passwords, and PINs.
package acopw

const (
	_indexBits = 6
	_indexMask = 1<<_indexBits - 1
)

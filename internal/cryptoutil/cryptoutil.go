// Package cryptoutil provides utility functions for cryptographic operations.
package cryptoutil

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"

	"git.sr.ht/~jamesponddotco/xstd-go/xerrors"
)

// ErrInvalidLength is returned when the length of the input is not valid.
const ErrInvalidLength xerrors.Error = "invalid length; must be >= 1"

// Int returns a uniform random value in [0, max). It works like rand.Int, but
// take a max parameter integer instead of a *big.Int.
func Int(reader io.Reader, max int) (int, error) {
	if max < 1 {
		return 0, fmt.Errorf("%w: %d", ErrInvalidLength, max)
	}

	result, err := rand.Int(reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, fmt.Errorf("failed to generate random number: %w", err)
	}

	return int(result.Int64()), nil
}

// RandomElement returns a random element from the given string slice.
func RandomElement(reader io.Reader, elements []string) (string, error) {
	if len(elements) < 1 {
		return "", fmt.Errorf("%w: %d", ErrInvalidLength, len(elements))
	}

	index, err := Int(reader, len(elements))
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	return elements[index], nil
}

// Package cryptoutil provides utility functions for cryptographic operations.
package cryptoutil

import (
	"encoding/binary"
	"fmt"
	"io"

	"git.sr.ht/~jamesponddotco/xstd-go/xerrors"
)

// ErrDivisionByZero is returned when a division by zero is attempted.
const ErrDivisionByZero xerrors.Error = "division by zero; input must be greater than zero"

// RandomIndex generates a cryptographically secure random index in the range
// [0, n-1]. It tries to eliminate bias by using the rejection method.
func RandomIndex(n int, randReader io.Reader) (int, error) {
	if n <= 0 {
		return 0, ErrDivisionByZero
	}

	var (
		max       = ^uint64(0)
		threshold = max - max%uint64(n)
		b         [8]byte
	)

	for {
		if _, err := io.ReadFull(randReader, b[:]); err != nil {
			return 0, fmt.Errorf("%w", err)
		}

		num := binary.BigEndian.Uint64(b[:])
		if num < threshold {
			return int(num % uint64(n)), nil
		}
	}
}

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
// [0, n-1]. It tries to eliminate bias by discarding values that would cause
// bias due to the modulo operation.
func RandomIndex(n int, randReader io.Reader) (int, error) {
	if n <= 0 {
		return 0, ErrDivisionByZero
	}

	var (
		max       = ^uint32(0) // Max value for uint32
		threshold = max - (max % uint32(n))
	)

	for {
		var r uint32

		if err := binary.Read(randReader, binary.LittleEndian, &r); err != nil {
			return 0, fmt.Errorf("%w", err)
		}

		if r < threshold {
			return int(r % uint32(n)), nil
		}
	}
}

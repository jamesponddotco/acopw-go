package cryptoutil_test

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"git.sr.ht/~jamesponddotco/acopw-go/internal/cryptoutil"
)

type errorReader struct{}

func (*errorReader) Read(_ []byte) (int, error) {
	return 0, errors.New("read error")
}

func TestInt(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		reader  io.Reader
		max     int
		wantErr bool
	}{
		{
			name:    "invalid length",
			reader:  nil,
			max:     0,
			wantErr: true,
		},
		{
			name:    "valid random number",
			reader:  bytes.NewReader([]byte{0, 1, 2, 3}),
			max:     4,
			wantErr: false,
		},
		{
			name:    "rand.Int error",
			reader:  &errorReader{},
			max:     4,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			_, err := cryptoutil.Int(tt.reader, tt.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("Int() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRandomElement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		reader   io.Reader
		elements []string
		wantErr  bool
	}{
		{
			name:     "invalid length",
			reader:   nil,
			elements: []string{},
			wantErr:  true,
		},
		{
			name:     "valid random element",
			reader:   bytes.NewReader([]byte{0, 1, 2, 3}),
			elements: []string{"apple", "banana", "cherry"},
			wantErr:  false,
		},
		{
			name:     "rand.Int error",
			reader:   &errorReader{},
			elements: []string{"apple", "banana", "cherry"},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			_, err := cryptoutil.RandomElement(tt.reader, tt.elements)
			if (err != nil) != tt.wantErr {
				t.Errorf("RandomElement() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

package output_test

import (
	"bytes"
	"errors"
	"io"
	"os"
	"strings"
	"testing"

	"git.sr.ht/~jamesponddotco/acopw-go/cmd/acopw/internal/output"
	"github.com/creack/pty"
)

type testWriter struct {
	buf bytes.Buffer
}

func (tw *testWriter) Write(p []byte) (n int, err error) {
	return tw.buf.Write(p)
}

func (tw *testWriter) String() string {
	return tw.buf.String()
}

func TestPassword(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		password string
		expected string
	}{
		{
			name:     "Password1",
			password: "password123",
			expected: "password123\n",
		},
		{
			name:     "Password2",
			password: "abc123!@#",
			expected: "abc123!@#\n",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Redirect the output.
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			output.Password(tt.password)

			// Restore the original output.
			w.Close()
			os.Stdout = oldStdout

			out, _ := io.ReadAll(r)
			if got := string(out); got != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, got)
			}
		})
	}
}

func TestError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		err      error
		expected string
	}{
		{
			name:     "Error1",
			err:      errors.New("Error 1"),
			expected: "[Error 1]\n",
		},
		{
			name:     "Error2",
			err:      errors.New("Error 2"),
			expected: "[Error 2]\n",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Redirect the output.
			oldStderr := os.Stderr
			r, w, _ := os.Pipe()
			os.Stderr = w

			output.Error(tt.err)

			// Restore the original output.
			w.Close()
			os.Stderr = oldStderr

			out, _ := io.ReadAll(r)
			if got := string(out); got != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, got)
			}
		})
	}
}

func TestPasswordTerminal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		password string
		expected string
	}{
		{
			name:     "Password1",
			password: "password123",
			expected: "password123",
		},
		{
			name:     "Password2",
			password: "abc123!@#",
			expected: "abc123!@#",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Redirect the output
			oldStdout := os.Stdout
			r, w, _ := pty.Open()
			os.Stdout = w

			output.Password(tt.password)

			// Restore the original output
			w.Close()
			os.Stdout = oldStdout

			out, _ := io.ReadAll(r)
			if got := strings.TrimSpace(string(out)); got != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, got)
			}
		})
	}
}

func TestErrorTerminal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		err      error
		expected string
	}{
		{
			name:     "Error1",
			err:      errors.New("Error 1"),
			expected: "[Error 1]",
		},
		{
			name:     "Error2",
			err:      errors.New("Error 2"),
			expected: "[Error 2]",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Redirect the output
			oldStderr := os.Stderr
			r, w, _ := pty.Open()
			os.Stderr = w

			output.Error(tt.err)

			// Restore the original output
			w.Close()
			os.Stderr = oldStderr

			out, _ := io.ReadAll(r)
			if got := strings.TrimSpace(string(out)); got != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, got)
			}
		})
	}
}

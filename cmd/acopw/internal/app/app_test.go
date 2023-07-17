package app_test

import (
	"strings"
	"testing"

	"git.sr.ht/~jamesponddotco/acopw-go/cmd/acopw/internal/app"
	"github.com/spf13/cobra"
)

func TestRun(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "Diceware",
			args:     []string{"diceware"},
			expected: "",
		},
		{
			name:     "PIN",
			args:     []string{"pin"},
			expected: "",
		},
		{
			name:     "Random",
			args:     []string{"random"},
			expected: "",
		},
		{
			name:     "UUID",
			args:     []string{"uuid"},
			expected: "",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			cmd := &cobra.Command{
				Use:               "test",
				CompletionOptions: cobra.CompletionOptions{HiddenDefaultCmd: true},
			}
			app.AddCommands(cmd)

			cmd.SetArgs(tt.args)
			err := cmd.Execute()

			if err != nil && !strings.Contains(err.Error(), tt.expected) {
				t.Errorf("Expected %q, got %q", tt.expected, err)
			}
		})
	}
}

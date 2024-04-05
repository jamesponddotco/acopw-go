package acopw_test

import (
	"regexp"
	"testing"

	"git.sr.ht/~jamesponddotco/acopw-go"
)

func TestUUID_Generate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want string
	}{
		{
			name: "Valid UUIDv4",
			want: "^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			u := &acopw.UUID{}
			got := u.Generate()

			matched, err := regexp.MatchString(tt.want, got)
			if err != nil {
				t.Fatalf("Failed to match regex: %v", err)
			}

			if !matched {
				t.Errorf("Generate() = %q, want match %q", got, tt.want)
			}
		})
	}
}

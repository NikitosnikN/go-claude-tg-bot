package utils

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEscapeMarkdown(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "No special characters",
			input: "Hello world",
			want:  "Hello world",
		},
		{
			name:  "Underscore",
			input: "Hello_world",
			want:  "Hello\\_world",
		},
		{
			name:  "Asterisk",
			input: "Hello*world",
			want:  "Hello\\*world",
		},
		{
			name:  "Backtick",
			input: "Hello`world`",
			want:  "Hello\\`world\\`",
		},
		{
			name:  "Square brackets",
			input: "Hello[world]",
			want:  "Hello\\[world\\]",
		},
		{
			name:  "Parentheses",
			input: "Hello(world)",
			want:  "Hello\\(world\\)",
		},
		{
			name:  "Tilde",
			input: "Hello~world",
			want:  "Hello\\~world",
		},
		{
			name:  "Greater than",
			input: "Hello>world",
			want:  "Hello\\>world",
		},
		{
			name:  "Hash",
			input: "Hello#world",
			want:  "Hello\\#world",
		},
		{
			name:  "Plus",
			input: "Hello+world",
			want:  "Hello\\+world",
		},
		{
			name:  "Minus",
			input: "Hello-world",
			want:  "Hello\\-world",
		},
		{
			name:  "Equal",
			input: "Hello=world",
			want:  "Hello\\=world",
		},
		{
			name:  "Vertical bar",
			input: "Hello|world",
			want:  "Hello\\|world",
		},
		{
			name:  "Curly braces",
			input: "Hello{world}",
			want:  "Hello\\{world\\}",
		},
		{
			name:  "Dot",
			input: "Hello.world",
			want:  "Hello\\.world",
		},
		{
			name:  "Exclamation mark",
			input: "Hello!world",
			want:  "Hello\\!world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EscapeMarkdown(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}

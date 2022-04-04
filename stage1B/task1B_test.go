package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunTask(t *testing.T) {
	tests := []struct {
		name            string
		input, expected []string
	}{
		{"test1", []string{"7", "81 22 31 44 41 78 98"}, []string{"686"}},
		{"test2", []string{"5", "81 14 88 2 22"}, []string{"308"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := strings.NewReader(strings.Join(tt.input, "\n") + "\n")
			expected := strings.Join(tt.expected, "\n") + "\n"

			var output bytes.Buffer
			RunTaskCustom(&output, input)
			got := output.String()

			if expected != got {
				t.Errorf("expected %s, got %s", expected, got)
			}
		})
	}
}

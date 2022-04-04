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
		{"test1", []string{"2", "6 2", "3 1"}, []string{"2", "1"}},
		{"test1", []string{"1", "1000 100"}, []string{"5"}},
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

package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunTask(t *testing.T) {
	tests := []struct {
		input, expected []string
	}{
		{[]string{"3", "3", "cba", "aaz", "2", "bc", "ab", "1", "a", "a"}, []string{"Yes", "No", "No"}},
		{[]string{"1", "3", "aac", "abb"}, []string{"Yes"}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			input := strings.NewReader(strings.Join(tt.input, "\n") + "\n")
			expected := strings.Join(tt.expected, "\n") + "\n"

			var output bytes.Buffer
			RunTaskCustom(&output, input)
			got := output.String()

			if expected != got {
				t.Errorf("expected %v, got %v, input %v", expected, got, input)
			}
		})
	}
}

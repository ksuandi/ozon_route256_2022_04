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
		{"test1", []string{"2", "1 2", "4 7", "3 6"}, []string{"1", "3", "0"}},
		{"test2", []string{"8", "9000000 10000000"}, []string{"0"}},
		{"test3", []string{"8", "600 789"}, []string{"0"}},
		{"test4", []string{"8", "4 7", "2 4", "9 8", "2 7", "4 6", "7 3", "9 14", "14 9", "5 2"},
			[]string{"3", "2", "0", "2", "0", "0", "0", "0", "0"}},
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

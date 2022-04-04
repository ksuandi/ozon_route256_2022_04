package main

import "testing"

func TestPalindrom(t *testing.T) {
	tests := []struct {
		Input, Palindrom string
	}{
		{Input: "1234", Palindrom: "1331"},
		{Input: "4567", Palindrom: "4664"},
		{Input: "45678", Palindrom: "45754"},
		{Input: "9876789", Palindrom: "9876789"},
		{Input: "998877778899", Palindrom: "998877778899"},
		{Input: "987654321", Palindrom: "987656789"},
		{Input: "987658321", Palindrom: "987666789"},
		{Input: "12", Palindrom: "22"},
		{Input: "90", Palindrom: "99"},
		{Input: "5", Palindrom: "5"},
		{Input: "10000", Palindrom: "10001"},
		{Input: "10020", Palindrom: "10101"},
		{Input: "1090", Palindrom: "1111"},
		{Input: "13999", Palindrom: "14041"},
	}

	for _, test := range tests {
		input := []byte(test.Input)
		pal := makeNearestPalindrom(input)
		pal2 := makeBiggerPalindrom(pal, input)

		if string(pal2) != test.Palindrom {
			t.Errorf("expected %v, got %v, input %v", test.Palindrom, string(pal2), test.Input)
		}

	}

}

package main

import (
	"fmt"
	"math"
	"strconv"
)

func makeNearestPalindrom(digits []byte) []byte {
	len := len(digits)
	if len == 1 {
		return digits
	}
	result := make([]byte, len)
	for i, j := 0, len-1; i <= j; i, j = i+1, j-1 {
		result[i] = digits[i]
		result[j] = digits[i]
	}
	return result
}

func makeBiggerPalindrom(pal, digits []byte) []byte {
	palValue, _ := strconv.Atoi(string(pal))
	digitsValue, _ := strconv.Atoi(string(digits))

	if palValue >= digitsValue {
		return pal
	}

	middle := int(math.Ceil(float64(len(pal))/2)) - 1
	//fmt.Println(middle, len(pal)/2)
	if string(pal[middle]) < "9" {
		pal[middle] += 1
		if middle < len(pal)/2 {
			pal[middle+1] += 1
		}
	} else {
		pal[middle] = byte('0')
		pal[middle-1] += 1
		return makeNearestPalindrom(pal)
	}
	return pal
}

func main() {
	var digits string
	fmt.Scanln(&digits)
	pal := makeNearestPalindrom([]byte(digits))
	pal = makeBiggerPalindrom(pal, []byte(digits))

	palValue, _ := strconv.Atoi(string(pal))
	digitsValue, _ := strconv.Atoi(string(digits))

	//fmt.Println(palValue, digitsValue)
	fmt.Println(palValue - digitsValue)
}

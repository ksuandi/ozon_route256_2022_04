package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Info struct {
	d, p string
}

func RunTask() {
	RunTaskCustom(os.Stdout, os.Stdin)
}

func RunTaskCustom(w io.Writer, r io.Reader) {
	input := scanInput(r)

	days, _ := strconv.Atoi(input[0])
	prices, _ := sliceToInt(strings.Split(input[1], " "))

	s := 1
	sum := 0

	for i := days; i > 0; i-- {
		maxPrice := findMax(prices)
		if maxPrice == prices[0] { // today max price - sell all
			sum += maxPrice * s
			s = 0
		}
		s += 1              // make more goods
		prices = prices[1:] // days passed
	}

	fmt.Fprintln(w, sum)
}

func scanInput(r io.Reader) []string {
	var input []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}
	return input
}

func sliceToInt(from []string) ([]int, error) {
	to := make([]int, len(from))
	for i, s := range from {
		v, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		to[i] = v
	}
	return to, nil
}

func findMax(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	RunTask()
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func RunTask() {
	RunTaskCustom(os.Stdout, os.Stdin)
}

type Network struct {
	n, r, k int
}

func RunTaskCustom(w io.Writer, r io.Reader) {
	input := scanInput(r)

	// parse
	// cnt, _ := strconv.Atoi(input[0])
	// if cnt > len(input)-1 {
	// 	panic(fmt.Sprintf("wrong input: %d, %d, %v", cnt, len(input), input))
	// }
	var networks []Network
	for i := 1; i < len(input); i++ {
		limits, _ := sliceToInt(strings.Split(input[i], " "))
		networks = append(networks, Network{limits[0], limits[1], 0})
	}

	// calc k
	for _, ntw := range networks {
		maxDistance := ntw.n / 2
		if ntw.r > maxDistance {
			ntw.k = 1
		} else {
			optimalK := float64(maxDistance) / float64(ntw.r)
			ntw.k = int(math.Ceil(optimalK))
		}

		// answer
		fmt.Fprintln(w, ntw.k)
	}
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

func main() {
	RunTask()
}

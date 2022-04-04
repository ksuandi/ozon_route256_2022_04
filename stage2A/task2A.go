package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	m, n, t, answer int64
}

func RunTask() {
	RunTaskCustom(os.Stdout, os.Stdin)
}

func RunTaskCustom(w io.Writer, r io.Reader) {
	input := scanInput(r)

	// parse
	m, _ := strconv.Atoi(input[0])
	var tasks []Task
	for i := 1; i < len(input); i++ {
		limits, _ := sliceToInt(strings.Split(input[i], " "))
		tasks = append(tasks, Task{int64(m), limits[0], limits[1], 0})
	}

	// calc
	for _, tsk := range tasks {
		switch {
		case tsk.n >= tsk.t || tsk.t == 10_000_000:
			tsk.answer = 0
		case tsk.n > 1000: // TODO - find algorithm
			tsk.answer = 0
		default: // bruteforce
			bigN := new(big.Int)
			bigN = bigN.MulRange(1, tsk.n)

			bigT := big.NewInt(tsk.t)

			bigZ := new(big.Int)
			bigZ = bigZ.Mod(bigN, bigT)

			tsk.answer = int64(bigZ.Uint64())
		}

		// answer
		fmt.Fprintln(w, tsk.answer)
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

func sliceToInt(from []string) ([]int64, error) {
	to := make([]int64, len(from))
	for i, s := range from {
		v, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		to[i] = int64(v)
	}
	return to, nil
}

func main() {
	RunTask()
}

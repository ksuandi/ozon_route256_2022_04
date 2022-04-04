package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func RunTask() {
	RunTaskCustom(os.Stdout, os.Stdin)
}

func RunTaskCustom(w io.Writer, r io.Reader) {
	maskNeeded := scanInput(r)

	sum := 0.0

	sum1 := 0.55
	sum2 := 11.0
	sum3 := 160.0

	thresholdForOne := int(math.Ceil(sum2 / sum1))
	thresholdForPackage := int(math.Ceil(sum3 / sum2))

	cargo1 := 0
	cargo2 := 0
	cargo3 := 0

	for maskNeeded > 0 {
		switch {
		case maskNeeded >= 24*thresholdForPackage:
			cargo3 += 1
			sum += sum3
			maskNeeded -= 24 * 16
		case maskNeeded >= 1*thresholdForOne:
			cargo2 += 1
			sum += sum2
			maskNeeded -= 24
		default:
			cargo1 += 1
			sum += sum1
			maskNeeded -= 1
		}
	}

	answer := fmt.Sprintf("%d %d %d", cargo1, cargo2, cargo3)
	fmt.Fprintln(w, answer)
}

func scanInput(r io.Reader) int {
	var input string
	fmt.Fscanln(r, &input)

	parsed, err := strconv.Atoi(input)
	if err != nil {
		return 0
	}

	return parsed
}

func main() {
	RunTask()
}

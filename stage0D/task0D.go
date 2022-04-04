package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

type Entry struct {
	part1, part2 string
	len          int
}

func RunTask() {
	RunTaskCustom(os.Stdout, os.Stdin)
}

func RunTaskCustom(w io.Writer, r io.Reader) {
	input := scanInput(r)
	for _, entry := range input {
		s1 := getMinString(entry.part1)
		s2 := getMaxString(entry.part2)
		answer := "Yes"
		if s1 >= s2 {
			answer = "No"
		}
		//fmt.Println(s1, s2)
		fmt.Fprintln(w, answer)
	}
}

func scanInput(r io.Reader) []Entry {
	var num int
	fmt.Fscanln(r, &num)

	input := make([]Entry, num)
	for i := 0; i < num; i++ {
		var (
			part1, part2 string
			len          int
		)
		fmt.Fscanln(r, &len)
		fmt.Fscanln(r, &part1)
		fmt.Fscanln(r, &part2)
		input[i] = Entry{part1, part2, len}
	}
	return input
}

func getMaxString(orig string) string {
	b := []byte(orig)
	sort.Slice(b, func(i, j int) bool { return b[i] > b[j] })
	return string(b)
}

func getMinString(orig string) string {
	b := []byte(orig)
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	return string(b)
}

func main() {
	RunTask()
}

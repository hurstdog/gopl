// Dup2 prints the count and text of lines that appear more than once in the
// input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	lineToFile := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, lineToFile)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, lineToFile)
			f.Close()
		}
	}
	// Interesting improvement to think about: How to print the lines sorted in
	// order of increasing (or decreasing) count?
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, lineToFile[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, lineToFile map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		t := input.Text()
		counts[t]++
		if !stringInSlice(f.Name(), lineToFile[t]) {
			lineToFile[t] = append(lineToFile[t], f.Name())
		}
	}
	// NOTE: Ignoring potential errors from input.Err()
}

func stringInSlice(str string, slice []string) bool {
	for _, val := range slice {
		if str == val {
			return true
		}
	}
	return false
}

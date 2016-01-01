// Comma prints the given numbers with commas added every three digits.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Printf("%s\n", comma(arg))
	}
}

// recursion!
func comma(s string) string {
	// If there's a dot in the number, only comma'-ize the part before the dot.
	if dot := strings.Index(s, "."); dot >= 0 {
		return comma(s[:dot]) + s[dot:]
	}
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

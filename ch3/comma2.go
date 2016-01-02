// Comma2 prints the given numbers with commas added every three digits,
// non-recursively.
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Printf("%s\n", comma(arg))
	}
}

// recursion!
func comma(s string) string {
	var buf bytes.Buffer
	i := len(s) % 3
	buf.WriteString(s[:i])
	for i < len(s) {
		if s[i] == '.' {
			buf.WriteString(s[i:])
			break
		}
		if i != 0 {
			buf.WriteString(",")
		}
		buf.WriteString(s[i : i+3])
		i += 3
	}
	return buf.String()
}

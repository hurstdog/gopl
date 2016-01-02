// Comma2 prints the given numbers with commas added every three digits,
// non-recursively.
package main

import (
	"bytes"
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
	var buf bytes.Buffer
	signpad := 0
	if len(s) > 0 && s[0] == '-' {
		signpad = 1
	}
	length := strings.Index(s, ".")
	if length < 0 {
		length = len(s)
	}
	length = length - signpad
	i := length%3 + signpad
	buf.WriteString(s[:i])
	for i < len(s) {
		if s[i] == '.' {
			buf.WriteString(s[i:])
			break
		}
		if i != signpad {
			buf.WriteString(",")
		}
		end := i + 3
		if end >= len(s) {
			buf.WriteString(s[i:])
			break
		}
		buf.WriteString(s[i:end])
		i += 3
	}
	return buf.String()
}

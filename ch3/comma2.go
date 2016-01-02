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
	dot := strings.Index(s, ".")
	l := dot
	if dot < 0 {
		l = len(s)
	}
	i := l % 3
	buf.WriteString(s[:i])
	for i < len(s) {
		if s[i] == '.' {
			buf.WriteString(s[i:])
			break
		}
		if i != 0 {
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

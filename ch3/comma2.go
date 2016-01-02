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
	// create a buffer.
	// Loop over the string backwards in groups of three, adding in a comma
	// between each.
	var c int
	var buf bytes.Buffer
	i := len(s) % 3
	buf.WriteString(s[:i] + ",")
	for ; i < len(s); i++ {
		if s[i] == '.' {
			buf.WriteString(s[i:])
			break
		}
		buf.WriteByte(s[i])
		c++
		if c > 2 {
			if i+1 < len(s) && s[i+1] != '.' {
				buf.WriteByte(',')
				c = 0
			}
		}
	}

	return buf.String()
}

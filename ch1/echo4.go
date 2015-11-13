// Echo4 prints its command line arguments, index and value.
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Printf("%d: %s\n", i, arg)
	}
}

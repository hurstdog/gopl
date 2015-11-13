// echo_threeways provides three functions for echoing text.
package ch1

import (
	"fmt"
	"os"
	"strings"
)

func EchoStringCat() {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func EchoPrintParts() {
	var sep string
	for _, arg := range os.Args[1:] {
		fmt.Print(sep, arg)
		sep = " "
	}
	fmt.Println()
}

func EchoStringJoin() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

// Reports if two words are anagrams of each other
package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: anagrams wordone wordtwo")
		os.Exit(1)
	}
	if anagrams(os.Args[1], os.Args[2]) {
		fmt.Println("Yes!")
	} else {
		fmt.Println("Nope!")
	}
}

func anagrams(s1 string, s2 string) bool {
	s1slice := strings.Split(s1, "")
	sort.Strings(s1slice)
	s1sorted := strings.Join(s1slice, "")
	s2slice := strings.Split(s2, "")
	sort.Strings(s2slice)
	s2sorted := strings.Join(s2slice, "")
	return s1sorted == s2sorted
}

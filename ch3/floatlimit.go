// Floatlimit prints the smallest integers that can't be accurately represented
// by a float.
package main

import (
	"fmt"
	"math"
)

func main() {
	var one32 float32 = 0
	for i := int32(0); i < math.MaxInt32; i++ {
		if one32 == one32+1 {
			fmt.Printf("float32: %f == %[1]f +1: %f!\n", one32, one32+1)
			break
		}
		one32 += 1
	}

	// For float64 cheat a little and start higher.
	var one64 float64 = 9.007199 * 10e14
	for i := int64(0); i < math.MaxInt64; i++ {
		if one64 == one64+1 {
			fmt.Printf("float64: %f == %[1]f +1: %f!\n", one64, one64+1)
			break
		}
		one64 += 1
	}
}

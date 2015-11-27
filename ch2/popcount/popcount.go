package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		// How does this work?
		// pc[0] = pc[0] + 0&1 = 0 + 0 = 0
		// pc[1] = pc[0] + 1&1 = 0 + 1 = 1
		// pc[2] = pc[1] + 2&1 = 1 + 0 = 1
		// pc[3] = pc[1] + 3&1 = 1 + 1 = 2
		// pc[4] = pc[2] + 4&1 = 1 + 0 = 1
		// pc[5] = pc[2] + 5&1 = 1 + 1 = 2
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCountLoop returns the population count (number of set bits) of x.
func PopCountLoop(x uint64) int {
	var res int
	for i := uint64(0); i < 8; i++ {
		res += int(pc[byte(x>>(i*8))])
	}
	return res
}

// PopCountShift returns the population count (number of set bits) of x.
func PopCountShift(x uint64) int {
	var res int
	for i := uint64(0); i < 64; i++ {
		res += int(x >> i & 1)
	}
	return res
}

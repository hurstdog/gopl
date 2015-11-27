package popcount

import "testing"

var testTable map[uint64]int = make(map[uint64]int)

func init() {
	testTable[0] = 0
	testTable[1] = 1
	testTable[2] = 1
	testTable[3] = 2
	testTable[4] = 1
	testTable[5] = 2
	testTable[6] = 2
}

func TestPopCount(t *testing.T) {
	for value, expected := range testTable {
		p := PopCount(value)
		if p != expected {
			t.Errorf("PopCount(%g): expected %g, got %g\n", value, expected, p)
		}
	}

}

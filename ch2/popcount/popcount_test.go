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
			t.Errorf("PopCount(%v): expected %v, got %v\n", value, expected, p)
		}
	}
}

func TestPopCountLoop(t *testing.T) {
	for value, expected := range testTable {
		p := PopCountLoop(value)
		if p != expected {
			t.Errorf("PopCountLoop(%v): expected %v, got %v\n", value, expected, p)
		}
	}
}

func TestPopCountShift(t *testing.T) {
	for value, expected := range testTable {
		p := PopCountShift(value)
		if p != expected {
			t.Errorf("PopCountShift(%v): expected %v, got %v\n", value, expected, p)
		}
	}
}

func TestPopCountShift2(t *testing.T) {
	for value, expected := range testTable {
		p := PopCountShift2(value)
		if p != expected {
			t.Errorf("PopCountShift2(%v): expected %v, got %v\n", value, expected, p)
		}
	}
}

func TestPopCountBitTrick(t *testing.T) {
	for value, expected := range testTable {
		p := PopCountBitTrick(value)
		if p != expected {
			t.Errorf("PopCountBitTrick(%v): expected %v, got %v\n", value, expected, p)
		}
	}
}

const testval = uint64(1231292910189123013)

func BenchmarkPopCountLoop(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PopCountLoop(testval)
	}
}

func BenchmarkPopCountRaw(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PopCount(testval)
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PopCountShift(testval)
	}
}

func BenchmarkPopCountShift2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PopCountShift2(testval)
	}
}

func BenchmarkPopCountBitTrick(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PopCountBitTrick(testval)
	}
}

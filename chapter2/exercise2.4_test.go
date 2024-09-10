// Compare pop count implementation looping and single expression
// byte wise table lookup

package popcount

import "testing"

// pc[i] is the population count of i.
// var pc [256]byte

func PopCountShift(x uint64) int {
	count := 0
	mask := uint64(1)

	for i := 0; i < 64; i++ {
		if x&mask > 0 {
			count++
		}
		x >>= 1
	}
	return count
}

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	bench(b, PopCountShift)

}

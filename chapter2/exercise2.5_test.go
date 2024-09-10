package popcount

import "testing"

func PopCountClearRightMostNonZero(x uint64) int {

	count := 0
	for x != 0 {
		x &= x - 1
		count++
	}
	return count
}

func BenchmarkRemoveRightMostNonZero(b *testing.B) {
	bench(b, PopCountClearRightMostNonZero)
}

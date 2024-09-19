// reverse n to left with one pass

package main

func rotateInThreePass(slice []int, n int) {
	reverseSlice(slice[:n])
	reverseSlice(slice[n:])
	reverseSlice(slice[:])
}
func rotateInOnePass(slice []int, n int) []int {
	slice = append(slice, reverseSlice(slice[0:n])...)
	return slice[n+1:]
}

func reverseSlice(slice []int) []int {
	for i, j := 0, len(slice)-1; j > i; j, i = j-1, i+1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

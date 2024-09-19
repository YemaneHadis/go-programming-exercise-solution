/***
	- Write a function that counts the number of bits that are different in two
	- SHA256 hashes (NB it is difference in bit not byte)
***/

package main

import (
	"crypto/sha256"
	"fmt"
)

func DifferentBits(h1, h2 *[sha256.Size]byte) int {
	// h1 and h2 are the sha256 representation of the two string
	count := 0
	for i := range h1 {
		for b := h1[i] ^ h2[i]; b != 0; b &= b - 1 {
			count++
		}
		/**
		 h2[i] ^ h2[i] - this will give us number of diffrent bit between thise two byte
		 if they are the same char this will be zero
		 else we will have one for different bits and 0 for same bits and the value is stored in b
		 b &= b-1 will remove 1 one bit from the byte  and we are incrmenting our  count
		**/
	}
	return count
}

func main() {
	s1 := "helloworld"
	s2 := "HELLOWORLD"
	h1 := sha256.Sum256([]byte(s1))
	h2 := sha256.Sum256([]byte(s2))
	fmt.Printf("%x\n", h1)
	fmt.Printf("Number of different bits between sha256 representation of %s %s are %d\n", s1, s2, DifferentBits(&h1, &h2))
}

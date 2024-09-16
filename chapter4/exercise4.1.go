package main

import (
	"fmt"
)

func main() {
	fmt.Printf("s1: %s h1: %X h1 type: %T\n", s1, h1, h1)
	fmt.Printf("s2: %s h2: %X h2 type: %T\n", s2, h2, h2)
	fmt.Printf("Number of different bits: %d\n", DifferentBits(h1, h2))
}

// count the number of bits set in x

// TODO READ THE FOLLOWING ARTICLE ABOUT BIT MANIPULATION
//https://graphics.stanford.edu/~seander/bithacks.html

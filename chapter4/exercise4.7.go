package main

import (
	"fmt"
	"unicode/utf8"
)

func reverseUTF8(str []byte) {
	size := len(str)
	for i := 0; i < len(str)/2; i++ {
		str[i], str[size-1-i] = str[size-1-i], str[i]
	}
}

// Reverse all the rune, and then entire slice
// credit https://github.com/torbiak

func main() {
	utf8String := []byte("Hello,  世界")
	fmt.Println(string(utf8String))
	for i := 0; i < len(utf8String); {
		_, size := utf8.DecodeRune(utf8String[i:])
		reverseUTF8(utf8String[i : i+size])
		i += size
	}
	fmt.Println(string(utf8String))

	reverseUTF8(utf8String)

	fmt.Println(string(utf8String))
}

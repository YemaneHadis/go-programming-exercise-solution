// IN-PLACE FUNCTION THAT SQUASHES EACH RUN OF ADJECENT UNICODE SPACES

package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func printChacter(strings []byte) {
	for i, w := 0, 0; i < len(strings); i += w {
		runeValue, width := utf8.DecodeRuneInString(string(strings[i:]))
		fmt.Printf("%#U starts at byte postion %d\n", runeValue, i)
		w = width
	}
}
func removeDuplicateSpace(strings []byte) []byte {
	space := false
	for i, w := 0, 0; i < len(strings); i += w {
		runeValue, width := utf8.DecodeRuneInString(string(strings[i:]))
		w = width
		if unicode.IsSpace(runeValue) {
			if space == true {
				// remove that space

				rightSide := strings[i+1:]
				strings = strings[:i]
				strings = append(strings, rightSide...)
			}
			{
				space = true
			}

		} else {
			space = false
		}

	}
	return strings
}

func main() {

	utf8String := "Hello,  世界"
	utf8Bytes := []byte(utf8String)
	fmt.Println("Before Removing")
	printChacter(utf8Bytes)
	utf8Bytes = removeDuplicateSpace(utf8Bytes)
	fmt.Println("After Removing")
	printChacter(utf8Bytes)

}

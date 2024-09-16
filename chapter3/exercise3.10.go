package main

import (
	"bytes"
	"fmt"
	"strings"
)

func reverseString(number string) string {
	runes := []rune(number)

	for i, j := 0, len(number)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)

}

func insertCommma(number string) string {
	var buff bytes.Buffer
	counter := 0
	// find the last index if . only inset command until that
	dotIndex := strings.LastIndex(number, ".")
	decimal := ""
	if dotIndex > 0 {
		decimal = number[dotIndex:]
		number = number[:dotIndex]
	}
	for i := len(number) - 1; i >= 0; i-- {
		if counter == 3 {
			buff.WriteByte(',')
			counter = 0
		}
		buff.WriteRune(rune(number[i]))
		counter++
	}

	return reverseString(buff.String()) + decimal
}

func main() {
	number := "1234.2112"
	fmt.Println(insertCommma(number))
	// fmt.Println(intsToString([]int{1, 3, 4}))
}

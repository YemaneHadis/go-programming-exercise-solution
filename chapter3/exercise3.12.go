package main

import "fmt"

func checkAnnaGram(string1 string, string2 string) bool {
	// change the two string into hashmap
	// and check for one element to be there in the other element
	str1Map := make(map[rune]int)

	str2Map := make(map[rune]int)

	for _, value := range string1 {
		str1Map[value]++
	}

	for _, value := range string2 {
		str2Map[value]++
	}

	for k, v := range str1Map {
		if str2Map[k] != v {
			return false
		}
	}

	for k, v := range str2Map {
		if str1Map[k] != v {
			return false
		}
	}

	return true

}

func main() {

	str1 := "anagram"
	str2 := "nagaram"

	fmt.Println(checkAnnaGram(str1, str2))
}

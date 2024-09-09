// modiy dup2 to print the names of all files in which each duplicate
// line occurs
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filesstring :=
		make(map[string][]string)

	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, filesstring)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2 : %v\n", err)
				continue
			}
			countLines(f, counts, filesstring)
			f.Close()

		}
	}

}

func countLines(f *os.File, counts map[string]int, filesstring map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if !in(f.Name(), filesstring[input.Text()]) {
			filesstring[input.Text()] = append(filesstring[input.Text()], f.Name())

		}
		counts[input.Text()]++

	}
	printMap(counts, filesstring)
	// Note egnoring the potential error from input.Err
}

func in(needle string, heyStack []string) bool {
	for _, v := range heyStack {
		if v == needle {
			return true
		}
	}
	return false
}

func printMap(counts map[string]int, filesstring map[string][]string) {
	for key, value := range counts {
		value2, exists := filesstring[key]
		if exists {
			fmt.Printf("Key: %s\n", key)
			fmt.Printf("filesstring1: %v\n", value)
			fmt.Printf("filesstring2: %v\n", value2)
		} else {
			fmt.Printf("Key %s does not exist in filesstring2\n", key)

		}
	}
}

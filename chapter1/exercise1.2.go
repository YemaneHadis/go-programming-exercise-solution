// print the index and value of each f its argumane

package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[0:] {
		fmt.Printf("%d \t %s \n", i, arg)
	}
}

// excerise 1.3 is similar to excerises in section 1.6 so skip for now

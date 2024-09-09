// modify the echo program to also print os.Args[0]
// the name of the command that invokes it

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print(strings.Join(os.Args[:], " "))
}

package main

import (
	"fmt"
	"os"
)

func main() {
	// Print index and value of each argument.
	for i, arg := range os.Args[1:] {
		fmt.Printf("[%2d] %s\n", i, arg)
	}
}

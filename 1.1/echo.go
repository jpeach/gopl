package main

import (
	"fmt"
	"os"
)

func main() {
	//	Also print the name of the invoked command ..
	fmt.Println(os.Args[:])
}

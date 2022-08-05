package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("Arguments value is", os.Args[1])
	} else {
		fmt.Println("No user arguments")
	}
}

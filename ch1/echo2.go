package main

import (
	"fmt"
	"os"
)

func main() {
	s := ""

	for _, arg := range os.Args[1:] {
		s += arg + " "
	}

	fmt.Printf("exec : %s\n", os.Args[0])
	fmt.Println(s)
}

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for j := len(os.Args) - 1; j >= 1; j-- {
		name := os.Args[j]
		for _, i := range name {
			z01.PrintRune(i)
		}

		z01.PrintRune('\n')
	}
}

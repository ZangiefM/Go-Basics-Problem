package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for j := 1; j < len(os.Args); j++ {
		name := os.Args[j]
		for _, i := range name {
			if i == '.' || i == '/' {
			} else {
				z01.PrintRune(i)
			}
		}

		z01.PrintRune('\n')
	}
}

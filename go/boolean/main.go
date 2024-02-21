package main

import (
	"os"

	"github.com/01-edu/z01"
)

type boolean bool

var (
	lengthOfArg int
	EvenMsg     string = "I have an even number of arguments"
	OddMsg      string = "I have an odd number of arguments"
)

func even(nbr int) int {
	if nbr%2 == 0 {
		return 1
	}
	return 0
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) boolean {
	if even(nbr) == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	lengthOfArg = len(os.Args) - 1
	if isEven(lengthOfArg) {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}

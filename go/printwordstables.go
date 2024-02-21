package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, word := range a {
		for i := 0; i < len(word); i++ {
			z01.PrintRune(rune(word[i]))
		}
		z01.PrintRune('\n')
	}
}

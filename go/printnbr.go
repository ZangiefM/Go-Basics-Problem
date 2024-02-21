package piscine

import "github.com/01-edu/z01"

func PrintNbr(a int) {
	numbers := []int{}
	index := 0
	minvalue := false
	if a < 0 {
		z01.PrintRune('-')
		if a ==   9223372036854775808 {
			a = a + 1
			minvalue = true
		}
		a = -1 * a
	} else if a == 0 {
		z01.PrintRune('0')
	}
	for a > 0 {
		numbers = append(numbers, a%10)
		a = a / 10
		index++
	}
	for i := len(numbers) - 1; i >= 0; i-- {
		if minvalue && i == 0 {
			z01.PrintRune(rune('8'))
		} else {
			z01.PrintRune(rune(48 + numbers[i]))
		}
	}
}

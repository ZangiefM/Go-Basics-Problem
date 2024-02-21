package piscine

func LastRune(s string) rune {
	a := []rune(s)
	i := len(s) - 1
	return rune(a[i])
}

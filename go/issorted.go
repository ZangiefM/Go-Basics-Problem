package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	result1 := true
	result2 := true
	n := len(a)
	for i := 0; i < n-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			result1 = false
			result1 = false

		}
	}
	for i := 0; i < n-1; i++ {
		if f(a[i], a[i+1]) < 0 {
			result2 = false

			result2 = false

		}
	}
	return result1 || result2
}

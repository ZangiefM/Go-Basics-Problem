package piscine

func Unmatch(arr []int) int {
	for _, res := range arr {
		test := 0
		for _, el := range arr {
			if el == res {
				test++
			}
		}
		if test == 1 || test%2 == 1 {
			return res
		}
	}
	return -1
}

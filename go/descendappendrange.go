package piscine

func DescendAppendRange(max, min int) []int {
	if min >= max {
		return []int{}
	}
	res := []int{}

	for i := max; i > min; i-- {
		res = append(res, i)
	}

	return res
}

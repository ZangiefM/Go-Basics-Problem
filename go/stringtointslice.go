package piscine

func StringToIntSlice(str string) []int {
	var bi []int
	for _, v := range str {
		bi = append(bi, int(v))
	}
	return bi
}

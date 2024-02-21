package piscine

func Compact(ptr *[]string) int {
	compacted := (*ptr)[0:0]
	for _, str := range *ptr {
		if str != "" {
			compacted = append(compacted, str)
		}
	}
	*ptr = compacted
	return len(*ptr)
}

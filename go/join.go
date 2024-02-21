package piscine

func Join(strs []string, sep string) string {
	var res string
	for i := range strs {
		if i < (len(strs) - 1) {
			res = res + strs[i]
			res = res + sep
		} else {
			res = res + strs[i]
		}
	}
	return res
}

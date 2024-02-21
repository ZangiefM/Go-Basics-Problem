package piscine

func IsNumeric(s string) bool {
	for _, i := range s {
		if i >= '0' && i <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}

package piscine

func IsLower(s string) bool {
	for _, i := range s {
		if i >= 'a' && i <= 'z' {
			continue
		} else {
			return false
		}
	}
	return true
}

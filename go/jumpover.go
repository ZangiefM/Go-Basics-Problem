package piscine

func JumpOver(str string) string {
	if len(str) == 0 || len(str) < 3 {
		return "\n"
	}

	result := ""
	for i, char := range str {
		// Include every third character
		if (i+1)%3 == 0 {
			result += string(char)
		}
	}

	return result + "\n"
}
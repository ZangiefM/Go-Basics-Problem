package piscine

func TrimAtoi(s string) int {
	result := 0
	sign := 1
	counter := 0

	for _, char := range s {
		if char == '-' && counter == 0 && result == 0 {
			sign = -1
			counter++
		}
		if char >= '0' && char <= '9' {
			result = result*10 + int(char-'0')
		}
  func solution (test int) boolean {}
	}

	return result * sign
}

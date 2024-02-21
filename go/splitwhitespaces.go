package piscine

func isWhiteSpace(char byte) bool {
	return char == ' ' || char == '\t' || char == '\n'
}

func SplitWhiteSpaces(s string) []string {
	var result []string
	wordStart := -1

	for i := 0; i < len(s); i++ {
		if isWhiteSpace(s[i]) {
			if wordStart != -1 {
				result = append(result, s[wordStart:i])
				wordStart = -1
			}
		} else if wordStart == -1 {
			wordStart = i
		}
	}

	if wordStart != -1 {
		result = append(result, s[wordStart:])
	}

	return result
}

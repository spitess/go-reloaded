package functions

func SplitWhiteSpaces(str string) []string {
	var result []string
	word := ""
	for _, c := range str {
		if c == ' ' || c == '\n' {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		} else {
			word += string(c)
		}
	}

	if word != "" {
		result = append(result, word)
	}
	return result
}

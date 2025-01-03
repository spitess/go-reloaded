package functions

func SplitWhiteSpaces(str string) []string {
	var result []string
	word := ""
	for _, c := range str {
		if c == ' ' {
			if word != "" {
				result = append(result, word) 
				word = ""
			}
		}
		if c != ' ' {
			word += string(c)
		}
	}

	if word != "" {
		result = append(result, word)
	}
	return result
}

package functions

func SplitWhiteSpaces(str string) []string {
	slice := []string{}
	word := ""
	for _, c := range str {
		if c == ' ' || c == '\t' || c == '\r' || c == '\v' || c == '\f' {
			if word != "" {
				slice = append(slice, word)
				word = ""
			}
		} else {
			word += string(c)
		}
	}

	if word != "" {
		slice = append(slice, word)
	}
	return slice
}

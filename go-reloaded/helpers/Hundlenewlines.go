package functions

func Hundlenewlines(str string) []string {
	slice := []string{}
	word := ""
	for _, c := range str {
		if c == '\n' {
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

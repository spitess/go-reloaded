package piscine

/// amine      amine

func trim(str string) string {
	res := ""
	for i, val := range str {
		if val == ' ' && i+1 < len(str) && str[i+1] == ' ' {
			continue
		}
		res += string(val)
	}
	return res
}

func Punctuations(s string) string {
	slice := []rune(trim(s))

	punc := ".,!?:;"
	for i, elm := range slice {
		for _, val := range punc {
			if (i+1 < len(slice)) && (elm == ' ' && slice[i+1] == val) {
				slice[i], slice[i+1] = slice[i+1], slice[i]
			}
		}
	}
	return string(slice)
}

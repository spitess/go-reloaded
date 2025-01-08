package functions

import (
	"strings"
)

func Punctuations(s string) string {
	slice := []rune(s)
// trimspace()
	check := ".,!?:;"

	for i, elm := range slice {
		for _, punc := range check { // check if we have a punc
			if (i+1 < len(slice)) && (elm == ' ' && slice[i+1] == punc) {
				slice[i], slice[i+1] = slice[i+1], slice[i]
			}
		}
	}
	// fmt.Println("slice -->" + string(slice))
	lastRes := strings.Join(SplitWhiteSpaces(string(slice)), " ")
	// fmt.Println("res-->" + lastRes)
	return lastRes
}

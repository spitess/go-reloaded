package functions

import (
	"strings"
)

func HundlPunctuations(s string) string {
	slice := []rune(s)
	check := ".,!?:;"

	for i, elm := range slice {
		for _, punc := range check {
			if (i+1 < len(slice)) && (elm == ' ' && slice[i+1] == punc) {
				slice[i], slice[i+1] = slice[i+1], slice[i]
			}
		}
	}
	lastRes := strings.Join(SplitWhiteSpaces(string(slice)), " ")
	return lastRes
}

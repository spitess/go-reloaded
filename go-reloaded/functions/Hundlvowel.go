package functions

import (
	r "goreload/helpers"
	"strings"
)

func HundlVowel(s string) string {
	slice := r.SplitWhiteSpaces(s)
	vowel := "aeiouhAEIOUH"
	for i, str := range slice {
		if str == "A" || str == "a" {
			if i+1 < len(slice) {
				word := slice[i+1]
				for _, vow := range vowel {
					if word[0] == byte(vow) {
						slice[i] = slice[i] + "n"
					}
				}
			}
		}
	}
	return strings.Join(slice, " ")
}

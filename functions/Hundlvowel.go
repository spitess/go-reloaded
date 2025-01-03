package functions

import (
	"strings"
)

func Hundlebowel(s string) string {
	res := strings.Split(s, " ")
	vowel := "aeiouhAEIOUH"
	for i, str := range res {
		if str == "A" || str == "a" {
			word := res[i+1]
			for _, vow := range vowel {
				if word[0] == byte(vow) {
					res[i] = "An"
				}
			}
		}
	}
	return strings.Join(res, " ")
}

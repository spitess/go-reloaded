package functions

import (
	"strings"
)

func Capitalize(s string) string {
	strmani := []rune(s)
	return strings.ToUpper(string(strmani[:1])) + strings.ToLower(string(strmani[1:]))

}

package functions

import (
	"strconv"
	"strings"
)

func HundlHexAndBin(s string) string {
	slice := SplitWhiteSpaces(s)
	flags := []string{"(hex)", "(bin)"}
	for i := 0; i < len(slice); i++ {
		if Contain(flags, slice[i]) && i > 0 {
			switch slice[i] {
			case "(hex)":
				value, err := strconv.ParseInt(slice[i-1], 16, 0)
				if err == nil {
					slice[i-1] = strconv.Itoa(int(value))
					slice[i] = ""
				}
			case "(bin)":
				value, err := strconv.ParseInt(slice[i-1], 2, 0)
				if err == nil {
					slice[i-1] = strconv.Itoa(int(value))
					slice[i] = ""
				}

			}
		}
	}
	res := strings.Join(slice, " ")
	str := Punctuations(res)
	return str
}

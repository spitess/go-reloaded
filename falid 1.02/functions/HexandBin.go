package functions

import (
	"strconv"
	"strings"

	H "goreload/helpers"
)

func HundlHexAndBin(s string) string {
	slice := H.SplitWhiteSpaces(s)
	flags := []string{"(hex)", "(bin)"}
	for i := 0; i < len(slice); i++ {
		flag := slice[i]
		if H.Contain(flags, flag) && i >= 0 {
			switch flag {
			case "(hex)":
				slice[i] = ""
				slice = H.RemoveEmptyStrings(slice)
				if i-1 >= 0 {
					value, err := strconv.ParseInt(slice[i-1], 16, 0)
					if err == nil {
						slice[i-1] = strconv.Itoa(int(value))
						i--
					}
				}
			case "(bin)":
				slice[i] = ""
				slice = H.RemoveEmptyStrings(slice)
				if i-1 >= 0 {
					value, err := strconv.ParseInt(slice[i-1], 2, 0)
					if err == nil {
						slice[i-1] = strconv.Itoa(int(value))
						i--
					}
				}
			}
			i--
		}
	}
	res := strings.Join(slice, " ")
	return res
}

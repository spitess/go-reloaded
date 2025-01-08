package functions

import (
	"fmt"
	"strconv"
	"strings"
)

func HundlHexAndBin(s string) string {

	slice := SplitWhiteSpaces(s)
	flags := []string{"(hex)", "(bin)"}
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(flags); j++ {
			if slice[i] == flags[j] {
				if flags[j] == "(hex)" {
					value, _ := strconv.ParseInt(slice[i-1], 16, 0)
					slice[i-1] = strconv.Itoa(int(value))
					slice[i] = ""
				} else if flags[j] == "(bin)" {
					value, _ := strconv.ParseInt(slice[i-1], 2, 0)
					fmt.Println("value 2---->", value)
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

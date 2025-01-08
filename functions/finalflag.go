package functions

import (
	"fmt"
	"strconv"
	"strings"
)

func Flg(s string) string {
	slice := SplitWhiteSpaces(s)
	flags := []string{"(cap)", "(up)", "(low)", "(cap,", "(low,", "(up,"}
	for i := 0; i < len(slice); i++ {
		if Contain(flags, slice[i]) && i > 0 {
			switch slice[i] {
			case "(cap)":
				slice[i-1] = Capitalize(slice[i-1])
				slice[i] = ""
			case "(up)":
				slice[i-1] = ToUpper(slice[i-1])
				slice[i] = ""
			case "(low)":
				slice[i-1] = ToLower(slice[i-1])
				slice[i] = ""
			case "(cap,":
				getvalue := slice[i+1]
				cleanvalue := strings.Trim(getvalue, ")")
				v, err := strconv.Atoi(cleanvalue)
				if err != nil || v < 0 {
					fmt.Println("Error: Invalid Flag:", cleanvalue, err)
				}
				if v >= 0 && v < i {
					start := i - v
					for k := start; k < i; k++ {
						slice[k] = Capitalize(slice[k])
						slice[i] = ""
					}
					slice[i+1] = ""
				}
			case "(low,":
				getvalue := slice[i+1]
				cleanvalue := strings.Trim(getvalue, ")")
				v, err := strconv.Atoi(cleanvalue)
				if err != nil || v < 0 {
					fmt.Println("Error: Invalid Flag:", cleanvalue, err)
				}
				if v >= 0 && v < i {
					start := i - v
					for x := start; x < i; x++ {
						slice[x] = ToLower(slice[x])
						slice[i] = ""
					}
					slice[i+1] = ""
				}
			case "(up,":
				getvalue := slice[i+1]
				cleanvalue := strings.Trim(getvalue, ")")
				v, err := strconv.Atoi(cleanvalue)
				if err != nil || v < 0 {
					fmt.Println("Error: Invalid Flag:", cleanvalue, err)
				}
				if v >= 0 && v < i {
					start := i - v
					for b := start; b < i; b++ {
						slice[b] = ToUpper(slice[b])
						slice[i] = ""
					}
					slice[i+1] = ""
				}
			}
		}
	}
	str := strings.Join(slice, " ")
	str = Punctuations(str)
	return str
}

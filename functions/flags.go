package functions

import (
	//"fmt"
	"strconv"
	"strings"
)

func Flags(s string) string {
	slice := SplitWhiteSpaces(s)
	// fmt.Println(slice)
	flags := []string{"(cap)", "(up)", "(low)", "(cap,", "(low,"}

	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(flags); j++ {
			if slice[i] == flags[j] {
				if i > 0 {
					if flags[j] == "(cap)" {
						slice[i-1] = Capitalize(slice[i-1])
						slice[i] = ""
					} else if flags[j] == "(up)" {
						slice[i-1] = ToUpper(slice[i-1])
						slice[i] = ""
					} else if flags[j] == "(low)" {
						slice[i-1] = ToLower(slice[i-1])
						slice[i] = ""
					} else if flags[j] == "(cap," {

						getvalue := slice[i+1]
						value := ""
						for _, char := range getvalue {
							if char >= '0' && char <= '9' {
								value += string(rune(char))
							}
							v, _ := strconv.Atoi(value)
							start := i - v
							for k := start; k < i; k++ {
								slice[k] = Capitalize(slice[k])
							}
						}
						slice[i], slice[i+1] = "", ""
					} else if flags[j] == "(low," {

						getvaluelow := slice[i+1]
						valuelow := ""
						for _, char := range getvaluelow {
							if char >= '0' && char <= '9' {
								valuelow += string(rune(char))
							}
						}
						v, _ := strconv.Atoi(valuelow)
						start := i - v
						for n := start; n < i; n++ {
							slice[n] = ToLower(slice[n])
						}
						slice[i], slice[i+1] = "", ""
					}
				}
			}
		}
		//	fmt.Println(slice[i])
	}
	str := strings.Join(slice, " ")
	res := strings.TrimSpace(str)
	return res
}

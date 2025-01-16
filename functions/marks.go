package functions

import (
	"strings"
)

func HundleMarks(s string) string {
	slice := []string{}
	count := 0
	for _, elm := range s {
		slice = append(slice, string(elm))
	}
	for i := 0; i <= len(slice)-1; i++ {
		if slice[i] == "'" {
			count++
			if count%2 != 1 {
				slice[i-1] = strings.Trim(slice[i-1], " ")
			} else {
				if i+1 < len(slice)-1 {
					slice[i+1] = strings.Trim(slice[i+1], " ")
				}
			}
		}
		if slice[i] == " " && slice[i+1] == " " {
			slice[i] = ""
		}
	}

	res := strings.Join(slice, "")
	return res
}

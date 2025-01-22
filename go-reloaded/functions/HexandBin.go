package functions
import (
	"fmt"
	r "goreload/helpers"
	"strconv"
	"strings"
)
func HundlHexAndBin(s string) string {
	slice := r.SplitWhiteSpaces(s)
	flags := []string{"(hex)", "(bin)"}
	for i := 0; i < len(slice); i++ {
		flag := slice[i]
		if r.Contain(flags, flag) && i >= 0 {
			switch flag {
			case "(hex)":
				slice[i] = ""
				slice = r.RemoveEmptyStrings(slice)
				if i-1 >= 0 {
					value, err := strconv.ParseInt(slice[i-1], 16, 0)
					if err == nil {
						slice[i-1] = strconv.Itoa(int(value))
						fmt.Printf("value2---> %T ", value)
						i--
					}
				}
			case "(bin)":
				slice[i] = ""
				slice = r.RemoveEmptyStrings(slice)
				if i-1 >= 0 {
					value, err := strconv.ParseInt(slice[i-1], 2, 0)
					if err == nil {
						slice[i-1] = strconv.Itoa(int(value))
						i--
					}
				}
			}
		}
	}
	res := strings.Join(slice, " ")
	return res
}

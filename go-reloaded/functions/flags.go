package functions

import (
	"strconv"
	"strings"
)

func Hundleflg(s string) string {
	slice := SplitWhiteSpaces(s)
	flags := []string{"(cap)", "(up)", "(low)", "(cap,", "(up,", "(low,"}
	for i := 0; i < len(slice); i++ {
		flag := slice[i]
		if Contain(flags, flag) && i <= len(slice) {
			switch flag {
			case "(cap)":

				if i >= 0 {
					slice[i] = ""
					slice = RemoveEmptyStrings(slice)
					if i-1 >= 0 {
						slice[i-1] = Capitalize(slice[i-1])
					}
					i--
				}
			case "(up)":
				if i >= 0 {
					slice[i] = ""
					slice = RemoveEmptyStrings(slice)
					if i-1 >= 0 {
						slice[i-1] = strings.ToUpper(slice[i-1])
					}
					i--
				}

			case "(low)":
				if i >= 0 {
					slice[i] = ""
					slice = RemoveEmptyStrings(slice)
					if i-1 >= 0 {
						slice[i-1] = strings.ToLower(slice[i-1])
					}
					i--
				}

			case "(cap,":
				if i+1 < len(slice) && strings.HasSuffix(slice[i+1], ")") {
					getflag := slice[i+1]
					cleanflag := strings.Trim(getflag, ")")
					v, err := strconv.Atoi(cleanflag)
					if err != nil {
						continue
					}
					if v >= 0 {
						start := i - v
						if start < 0 {
							start = 0
						}
						for a := start; a < i; a++ {
							slice[a] = Capitalize(slice[a])
						}
						slice[i] = ""
						slice[i+1] = ""
						slice = RemoveEmptyStrings(slice)
						i--
					}
				}

			case "(up,":
				if i+1 < len(slice) && strings.HasSuffix(slice[i+1], ")") {
					getflag := slice[i+1]
					cleanflag := strings.Trim(getflag, ")")
					v, err := strconv.Atoi(cleanflag)
					if err != nil {
						continue
					}
					if v >= 0 {
						start := i - v
						if start < 0 {
							start = 0
						}

						for b := start; b < i; b++ {
							slice[b] = strings.ToUpper(slice[b])
						}
						slice[i] = ""
						slice[i+1] = ""
						slice = RemoveEmptyStrings(slice)
						i--
					}

				}

			case "(low,":
				if i+1 < len(slice) && strings.HasSuffix(slice[i+1], ")") {
					getflag := slice[i+1]
					cleanflag := strings.Trim(getflag, ")")
					v, err := strconv.Atoi(cleanflag)
					if err != nil {
						continue
					}
					if v >= 0 {
						start := i - v
						if start < 0 {
							start = 0
						}

						for c := start; c < i; c++ {
							slice[c] = strings.ToLower(slice[c])
						}
						slice[i] = ""
						slice[i+1] = ""
						slice = RemoveEmptyStrings(slice)
						i--
					}

				}

			}
		}

	}
	res := strings.Join(slice, " ")
	return res
}

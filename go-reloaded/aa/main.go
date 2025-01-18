package main

import "fmt"

func ee(str string) []string{      //"ghdiha            ..            ," ===== "ghdiha  .. ,"
	//puncstr := ".,!?:;"
	str1 := ""
	//p := 0
	aa := []string{}
	for i:=0; i<len(str); i++ {
		if string(str[i]) != " " {
			str1 += string(str[i])
		}else {
			if str1 != "" {
				aa = append(aa, str1)
				str1 = ""
			}
		}
	}
	if str1 != "" {
		aa = append(aa, str1)
		str1 = ""
	}
return aa
}
func main() {
	fmt.Print(ee("ghdiha            ..            , "))
}
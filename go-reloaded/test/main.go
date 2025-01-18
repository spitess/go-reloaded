package main

import (
	"fmt"
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

func SplitWhiteSpaces(str string) []string {
	var result []string
	word := ""
	for _, c := range str {
		if c == ' ' {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		} else {
			word += string(c)
		}
	}

	if word != "" {
		result = append(result, word)
	}
	return result
}

func HandlePunctuations(s string) string {
	slice := []rune(s)
	punc := ".,!?:;"
	fmt.Println("len(slice)-1----->", len(slice)-1)
	fmt.Println("len(slice)----->", len(slice))
	for i := 0; i < len(slice)-1; i++ {
		for _, v := range punc {
			if slice[i] == ' ' && slice[i+1] == v {
				slice = append(slice[:i], slice[i+1:]...)
			}
		}
	}

	return string(slice)
}

func main() {
	input := "aa , .. ; "

	res := HandlePunctuations(input)

	fmt.Println(res)
}

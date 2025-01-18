package main

import (
	"fmt"
	"strconv"

	// "strings"

	piscine "piscine/functions"
)

func Flags(s string) []string {
	slice := piscine.SplitWhiteSpaces(s)
	// fmt.Println(slice)
	flags := []string{"(cap)", "(up)", "low", "(cap,", "(low,"}

	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(flags); j++ {
			if slice[i] == flags[j] {
				if i > 0 {
					if flags[j] == "(cap)" {
						slice[i-1] = piscine.Capitalize(slice[i-1])
					} else if flags[j] == "(up)" {
						slice[i-1] = piscine.ToUpper(slice[i-1])
					} else if flags[j] == "(low)" {
						slice[i-1] = piscine.ToLower(slice[i-1])
					} else if flags[j] == "(cap," {
						getvalue := slice[i+1]
						value := ""
						for _, char := range getvalue {
							if char >= '0' && char <= '9' {
								value += string(rune(char))
							}
						}
						v, _ := strconv.Atoi(value)

						start := i - v
						fmt.Println("value of i ----->", i)
						fmt.Println("start---->", start)
						for k := start; k < i; k++ {
							slice[k] = piscine.Capitalize(slice[k])
						}
					} else if flags[j] == "(low," {
						getvaluelow := slice[i+1]
						valuelow := ""
						for _, charlow := range getvaluelow {
							if charlow >= '0' && charlow <= 9 {
								valuelow += string(rune(charlow))
							}
							x, _ := strconv.Atoi(valuelow)
							for b := x; b < i; b++ {
								slice[b] = piscine.ToLower(slice[b])
							}
						}
					}
				}
			}
		}
	}
	return slice
}
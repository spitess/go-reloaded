package main

import (
	"fmt"
	"os"
	piscine "piscine/functions"
)

func main() {
	Args := os.Args[1:]
	if len(Args) != 2 {
		fmt.Println("Err: You Must To Add Two Arguments")
		return
	}
	content, err := os.ReadFile(Args[0])
	if err != nil {
		fmt.Println(err.Error())
	}
	if string(content) == "" {
		fmt.Print("the file is empty")
	}
	fllags := piscine.Flg(string(content))
	punc := piscine.Punctuations(fllags)
	vowel := piscine.HundlVowel(punc)
	marks := piscine.HundleMarks(vowel)
	hex := piscine.HundlHexAndBin(marks)
	os.WriteFile("result.txt", []byte(hex), 0o664)
	fmt.Println("last---->", hex)
}

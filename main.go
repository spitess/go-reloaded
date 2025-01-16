package main

import (
	"fmt"
	"os"

	goreload "goreload/functions"
)

func main() {
	Args := os.Args[1:]
	if len(Args) != 2 {
		fmt.Println("Error : You Must To Add Two Arguments")
		return

	}
	if Args[0] != "sample.txt" || Args[1] != "result.txt" {
		fmt.Println("Error : The file doesn't exist.")
		return
	}

	content, err := os.ReadFile(Args[0])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if string(content) == "" {
		fmt.Print("Error : The file is empty")
		return
	}
	fllags := goreload.Hundleflg(string(content))
	hexandbin := goreload.HundlHexAndBin(fllags)
	punc := goreload.HundlPunctuations(hexandbin)
	marks := goreload.HundleMarks(punc)
	vowel := goreload.HundlVowel(marks)
	os.WriteFile("result.txt", []byte(vowel), 0o664)
	fmt.Println("result---->", vowel)
}

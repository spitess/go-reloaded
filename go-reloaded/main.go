package main

import (
	"fmt"
	"os"
	"strings"

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
	result := []string{}

	lines := strings.Split(string(content), "\n") // split with new line
	//  line1 line2  line3
	for _, line := range lines { // range on lines to read line by line [aaaaa,bbbbbb,cccccc]
		fllags := goreload.Hundleflg(line)
		hexandbin := goreload.HundlHexAndBin(fllags)
		punc := goreload.HundlPunctuations(hexandbin)
		marks := goreload.HundleMarks(punc)
		vowel := goreload.HundlVowel(marks)
		result = append(result, vowel)
	}

	os.WriteFile("result.txt", []byte(strings.Join(result, "\n")), 0o664) // [aaaaa\n,bbbbbb\n,cccccc\n]
	fmt.Println("result---->", result)
}

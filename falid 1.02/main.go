package main

import (
	"fmt"
	"os"
	str "strings"

	goreload "goreload/functions"
)

func main() {
	Args := os.Args[1:]
	if len(Args) != 2 {
		fmt.Println("Error : You Must To Add Two Arguments")
		return

	}

	content, err := os.ReadFile(Args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	if Args[1] != "result.txt" {
		fmt.Println("Error : The File Doesn't Exist.", Args[1])
		return
	}

	if string(content) == "" {
		fmt.Print("Error : The File Is Empty")
		return
	}

	Resfinal := []string{}
	slicelines := str.SplitAfter(string(content), "\n")
	for _, line := range slicelines {
		fllags := goreload.Hundleflg(line)
		hexandbin := goreload.HundlHexAndBin(fllags)
		vowel := goreload.HundlVowel(hexandbin)
		marks := goreload.HundleMarks(vowel)
		punc := goreload.HundlPunctuations(marks)
		Resfinal = append(Resfinal, punc)
	}
	newslice := []byte{}
	for _, v := range Resfinal {
		newslice = append(newslice, []byte(v)...)
	}
	os.WriteFile("result.txt", newslice, 0o664)
}

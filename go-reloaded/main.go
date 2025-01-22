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
		fmt.Println(err.Error())
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
	slicelines := str.Split(string(content), "\n")
	for _, line := range slicelines {
		fllags := goreload.Hundleflg(line)
		hexandbin := goreload.HundlHexAndBin(fllags)
		punc := goreload.HundlPunctuations(hexandbin)
		marks := goreload.HundleMarks(punc)
		vowel := goreload.HundlVowel(marks)
		Resfinal = append(Resfinal, vowel)
	}

	os.WriteFile("result.txt", []byte(str.Join(Resfinal, "\n")), 0o664)
	// fmt.Println("Lines---->", Resfinal)
}

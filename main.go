package main

import (
	"fmt"
	"os"

	piscine "piscine/functions"
)

func main() {
	Args := os.Args[1:]

	if len(Args) != 2 { // check if we have 2 args
		fmt.Println("Err: You Must To Add Two Arguments")
		return
	}
	content, err := os.ReadFile(Args[0]) // read the file
	if err != nil {                      // check if we have error
		fmt.Println(err.Error())
	}
	if string(content) == "" { // check if the file is empty
		fmt.Print("the file is empty")
	}

	// output := piscine.Punctuations(string(content))
	fllags := piscine.Flags(string(content))
	punc := piscine.Punctuations(fllags)

	os.WriteFile("result.txt", []byte(punc), 0o664) // copy and write the content to another file
	fmt.Println("punc---->", punc)                  // debug output
}

package main

import (
	"fmt"
	"os"
	"strings"

	parsing "goreload/functions/AllFunc"
)

func main() {
	Args := os.Args[1:]
	if len(Args) != 2 {
		fmt.Println("Error : You Must To Add Two Arguments")
		return

	}
	content, err := os.ReadFile(Args[0])
	switch {
	case err != nil:
		fmt.Println(err)
		return
	case !strings.HasSuffix(Args[1], ".txt"):
		fmt.Println("Error: The output file must have a .txt extension.")
		return
	case string(content) == "":
		fmt.Println("Please provide valid content.")
		return

	}
	
	parsing := parsing.AllFunc(string(content))
	os.WriteFile(Args[1], []byte(parsing), 0o664)
}

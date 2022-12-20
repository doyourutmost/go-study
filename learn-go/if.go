package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "abc.txt"
	if content, err := os.ReadFile(filename); err != nil {
		fmt.Println("cannot print file content: ", err)
	} else {
		fmt.Println(string(content))
	}
}

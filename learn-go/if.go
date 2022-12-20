package main

import (
	"fmt"
	"os"
)

// if的条件里可以赋值
// if的条件里赋值的变量作用域就在这个if语句里
func main() {
	filename := "abc.txt"
	if content, err := os.ReadFile(filename); err != nil {
		fmt.Println("cannot print file content: ", err)
	} else {
		fmt.Println(string(content))
	}
}

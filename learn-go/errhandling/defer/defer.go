package main

import (
	"bufio"
	"fmt"
	"learn-go/func/fib"
	"os"
)

func tryDefer() {
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//panic("Error occurred")
	//return
	//fmt.Println(3)
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func writeFile(filename string) {
	//file, err := os.Create(filename)
	file, err := os.OpenFile(
		filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		//panic(err)

		//fmt.Println("file already exists")
		//return

		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	writer := bufio.NewWriter(file)

	defer func(writer *bufio.Writer) {
		err := writer.Flush()
		if err != nil {
			panic(err)
		}
	}(writer)
	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		_, err := fmt.Fprintln(writer, f())
		if err != nil {
			return
		}
	}
}

func main() {
	//tryDefer()
	writeFile("fib.txt")
}

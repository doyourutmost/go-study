package main

import (
	"fmt"
	"learn-go/interface/test"
)

func getRetriever() retriever {
	return test.Retriever{}
}

type retriever interface {
	Get(string) string
}

func main() {
	var r retriever = getRetriever()
	fmt.Println(r.Get("https://www.imooc.com"))
}

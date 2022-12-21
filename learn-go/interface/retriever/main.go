package main

import "fmt"

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("https://www.imooc.com")
}
func main() {
	var r Retriever
	fmt.Println(download(r))
}

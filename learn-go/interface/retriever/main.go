package main

import (
	"fmt"
	"learn-go/interface/retriever/mock"
	"learn-go/interface/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(rp RetrieverPoster) string {
	rp.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return rp.Get(url)
}

const url = "https://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		})

}
func main() {
	var r Retriever
	pr := &mock.Retriever{Contents: "fake news"}
	inspect(pr)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)

	// Type assertion
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}
	//fmt.Println(download(r))
	fmt.Println(session(pr))

}

func inspect(r Retriever) {
	fmt.Println("Inspecting :", r)
	fmt.Printf(" > %T %v\n", r, r)
	fmt.Printf(" > type switch")

	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents: ", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)
	}
}

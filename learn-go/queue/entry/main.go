package main

import (
	"fmt"
	"learn-go/queue"
)

func main() {
	q := queue.Queue{3}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	q.Push("abc")
	fmt.Println(q.Pop())

}

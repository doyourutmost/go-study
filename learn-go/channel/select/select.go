package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() (out chan int) {
	out = make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return
}

func worker(id int, c chan int) {
	for val := range c {
		time.Sleep(1 * time.Second)
		fmt.Printf("Worker %d received %d\n", id, val)
	}
}
func createWorker(id int) (out chan int) {
	out = make(chan int)
	go worker(id, out)
	return
}

func main() {
	var c1, c2 chan int = generator(), generator() // c1,c2 = nil
	var worker = createWorker(0)

	n := 0
	//hasValue := false
	var values []int

	tm := time.After(10 * time.Second) // 10 秒后返回一个chan
	tick := time.Tick(time.Second)     // 每秒返回一个 chan
	for {
		var activeWorker chan int
		var activeValue int

		//if hasValue {
		//	activeWorker = worker
		//}
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n = <-c1:
			//hasValue = true
			values = append(values, n)
		case n = <-c2:
			//hasValue = true
			values = append(values, n)
		case activeWorker <- activeValue:
			//hasValue = false
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("queue len = ", len(values))

		case endTime := <-tm:
			fmt.Println("bye: ", endTime)
			return
		}
	}
}

package main

import (
	"fmt"
	"sync"
	"time"
)

type atmoicInt struct {
	value int
	lock  sync.Mutex
}

func (a *atmoicInt) increment() {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.value++
}
func (a *atmoicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}
func main() {
	var a atmoicInt
	a.increment()
	go a.increment()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}

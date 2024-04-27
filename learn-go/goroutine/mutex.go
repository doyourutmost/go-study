// Package main
// @title: mutex
// @description: 描述
// @author: CaoChao
// @date: 2024-04-22
package main

import (
	"fmt"
	"sync"
)

var (
	sum int
	wg  sync.WaitGroup
	mu  sync.Mutex
)

func add() {
	for i := 0; i < 10000000; i++ {
		mu.Lock()
		sum = sum + 1
		mu.Unlock()
	}
}

func sub() {
	for i := 0; i < 10000000; i++ {
		mu.Lock()
		sum = sum - 1
		mu.Unlock()
	}
}

func main() {
	wg.Add(2)
	go func() {
		defer wg.Done()
		add()
	}()
	go func() {
		defer wg.Done()
		sub()
	}()
	wg.Wait()
	fmt.Println("sum: ", sum)
}

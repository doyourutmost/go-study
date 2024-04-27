// Package main
// @title: mutex
// @description: 描述
// @author: CaoChao
// @date: 2024-04-22
package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg1 sync.WaitGroup
	rw  sync.RWMutex
)

func read() {
	wg1.Add(1)
	defer wg1.Done()
	rw.RLock()
	defer rw.RUnlock()
	fmt.Println("开始读取数据")
	time.Sleep(time.Second)
	fmt.Println("读取数据完毕")
}

func write() {
	wg1.Add(1)
	defer wg1.Done()
	rw.Lock()
	defer rw.Unlock()
	fmt.Println("开始写入数据")
	time.Sleep(time.Second * 3)
	fmt.Println("写入数据完毕")
}

func main() {
	for i := 0; i < 100; i++ {
		go read()
	}
	for i := 0; i < 10; i++ {
		go write()
	}
	wg1.Wait()
}

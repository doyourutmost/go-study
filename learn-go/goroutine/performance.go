// Package main
// @title: performance
// @description: 描述
// @author: CaoChao
// @date: 2024-04-22
package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 1000000; i++ {
		go func(i int) {
			for {
				fmt.Println(i)
				time.Sleep(time.Second)
			}
		}(i)
	}
	time.Sleep(30 * time.Second)
}

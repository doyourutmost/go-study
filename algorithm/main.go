package main

import (
	"algorithm/sort"
	"math/rand"
	"sync"
	"time"
)

func task(wg *sync.WaitGroup, s sort.SortFunc, arr []int) {
	copyArr := make([]int, len(arr))
	copy(copyArr, arr)
	sort.ConsumeSortTime(s, copyArr)
	defer wg.Done()
}
func main() {
	rand.Seed(time.Now().UnixNano())
	arr := rand.Perm(1e8)
	var wg sync.WaitGroup
	wg.Add(3)
	// go task(&wg, sort.SelectionSort, arr)
	// go task(&wg, sort.BubbleSort, arr)
	// go task(&wg, sort.InsertionSort, arr)
	go task(&wg, sort.ShellSort, arr)
	go task(&wg, sort.MergeSort, arr)
	go task(&wg, sort.QuickSort, arr)
	wg.Wait()
}

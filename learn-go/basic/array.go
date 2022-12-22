package main

import (
	"fmt"
)

func init() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 4}
	arr3 := [...]int{23, 4, 123, 2}
	var grid [4][5]bool
	fmt.Println(arr1, arr2, arr3, grid)
}
func main() {
}

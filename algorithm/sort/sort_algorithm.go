package sort

import (
	"fmt"
	"reflect"
	"regexp"
	"runtime"
	"time"
)

type SortFunc func([]int) []int

// BubbleSort
// 冒泡排序
// 稳定性：稳定
// 时间复杂度 ：最佳：O(n) ，最差：O(n2)， 平均：O(n2)
// 空间复杂度 ：O(1)
// 排序方式 ：In-place
// <p>
// 算法步骤:
// （1）比较相邻的元素。如果第一个比第二个大，就交换它们两个；
// （2）对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对，这样在最后的元素应该会是最大的数；
// （3）针对所有的元素重复以上的步骤，除了最后一个；
// （4）重复步骤 1~3，直到排序完成。
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		flag := false
		for j := 1; j < len(arr)-i; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return arr
}

// SelectionSort
// 选择排序
// 稳定性：不稳定
// 时间复杂度 ：最佳：O(n2) ，最差：O(n2)， 平均：O(n2)
// 空间复杂度 ：O(1)
// 排序方式 ：In-place
// 算法步骤:
// （1）首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
// （2）再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
// （3）重复第 2 步，直到所有元素均排序完毕。
func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		if minIndex != i {
			arr[minIndex], arr[i] = arr[i], arr[minIndex]
		}
	}
	return arr
}

// InsertionSort
// 插入排序
// 稳定性：稳定
// 时间复杂度 ：最佳：O(n) ，最差：O(n2)， 平均：O(n2)
// 空间复杂度 ：O(1)
// 排序方式 ：In-place
// 算法步骤：
// （1）从第一个元素开始，该元素可以认为已经被排序；
// （2）取出下一个元素，在已经排序的元素序列中从后向前扫描；
// （3）如果该元素（已排序）大于新元素，将该元素移到下一位置；
// （4）重复步骤 3，直到找到已排序的元素小于或者等于新元素的位置；
// （5）将新元素插入到该位置后；
// （6）重复步骤 2~5。
func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		curVal := arr[i]
		preIndex := i
		for ; preIndex >= 1 && arr[preIndex-1] > curVal; preIndex-- {
			arr[preIndex] = arr[preIndex-1]
		}
		arr[preIndex] = curVal
	}
	return arr
}

func checkSorted(arr []int) bool {
	if len(arr) <= 1 {
		return true
	}
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			panic("数组无序")
		}
	}
	return true
}

func ConsumeSortTime(s SortFunc, arr []int) {
	start := time.Now() // 记录开始时间
	orderSort := s(arr)
	elapsed := time.Since(start) // 记录结束时间
	checkSorted(orderSort)
	fmt.Printf("[%s]排序算法耗时：%s\n", getFunctionName(s), elapsed.String())
}

func getFunctionName(i interface{}) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	regex := regexp.MustCompile(`[^.]+$`)
	funcName := regex.FindString(fullName)
	return funcName
}

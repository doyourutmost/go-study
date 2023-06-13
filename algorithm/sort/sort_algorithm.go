package sort

import (
	"fmt"
	"math/rand"
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

// ShellSort
// 希尔排序
// 稳定性：不稳定
// 时间复杂度 ：最佳：O(nlogn)， 最差：O(n2) 平均：O(nlogn)
// 空间复杂度 ：O(1)
// 最差：O(n2)列子：5,1,6,2,7,3,8,4
// 排序方式 ：In-place
// 算法步骤：
// （1）选择一个增量序列 {t1, t2, …, tk}，其中 (ti>tj, i<j, tk=1)；
// （2）按增量序列个数 k，对序列进行 k 趟排序；
// （3）每趟排序，根据对应的增量 t，将待排序列分割成若干长度为 m 的子序列，分别对各子表进行直接插入排序。仅增量因子为 1 时，整个序列作为一个表来处理，表长度即为整个序列的长度。
func ShellSort(arr []int) []int {
	incFactor := 2
	for group := len(arr) / incFactor; group > 0; group /= incFactor {
		for i := 1; i < len(arr); i++ {
			curVal := arr[i]
			preIndex := i
			for ; preIndex-group >= 0 && arr[preIndex-group] > curVal; preIndex -= group {
				arr[preIndex] = arr[preIndex-group]
			}
			arr[preIndex] = curVal
		}
	}
	return arr
}

// MergeSort
// 归并排序
// 稳定性：稳定
// 时间复杂度 ：最佳：O(nlogn)， 最差：(nlogn) 平均：(nlogn)
// 空间复杂度 ：O(n)
// 排序方式 ：Out-place
// 算法步骤：
// 归并排序算法是一个递归过程，边界条件为当输入序列仅有一个元素时，直接返回，具体过程如下：
// （1）如果输入内只有一个元素，则直接返回，否则将长度为 n 的输入序列分成两个长度为 n/2 的子序列；
// （2）分别对这两个子序列进行归并排序，使子序列变为有序状态；
// （3）设定两个指针，分别指向两个已经排序子序列的起始位置；
// （4）比较两个指针所指向的元素，选择相对小的元素放入到合并空间（用于存放排序结果），并移动指针到下一位置；
// （5）重复步骤 3 ~ 4 直到某一指针达到序列尾；
// （6）将另一序列剩下的所有元素直接复制到合并序列尾。
func MergeSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	arr1 := arr[0 : length/2]
	arr2 := arr[length/2+1:]

	return mergeSort(MergeSort(arr1), MergeSort(arr2))
}

func mergeSort(arr1 []int, arr2 []int) (arr []int) {
	arr = make([]int, 0, len(arr1)+len(arr2))
	index1 := 0
	index2 := 0
	for index1 < len(arr1) && index2 < len(arr2) {
		if arr1[index1] < arr2[index2] {
			arr = append(arr, arr1[index1])
			index1++
		} else {
			arr = append(arr, arr2[index2])
			index2++
		}
	}
	if index1 < len(arr1) {
		arr = append(arr, arr1[index1:]...)
	} else if index2 < len(arr2) {
		arr = append(arr, arr2[index2:]...)
	}
	return
}

// QuickSort
// 快速排序
// 稳定性：不稳定
// 时间复杂度 ：最佳：O(nlogn)， 最差：O(nlogn)，平均：O(nlogn)
// 空间复杂度 ：O(nlogn)
// 排序方式 ：In-place
// 算法步骤：
// 快速排序算法是一个递归过程，边界条件为当输入序列仅有一个元素时，直接返回，具体过程如下：
// （1）从序列中随机挑出一个元素，做为 “基准”(pivot)；
// （2）重新排列序列，将所有比基准值小的元素摆放在基准前面，所有比基准值大的摆在基准的后面（相同的数可以到任一边）。在这个操作结束之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
// （3）递归地把小于基准值元素的子序列和大于基准值元素的子序列进行快速排序。
func QuickSort(arr []int) []int {
	quickSort(arr, 0, len(arr)-1)
	return arr
}

// quickSort 左右指针交换法
// func quickSort(arr []int, startIndex, endIndex int) {
// 	left := startIndex
// 	right := endIndex
//
// 	indexRange := endIndex - startIndex + 1
// 	randomIndex := rand.Intn(indexRange) + startIndex
// 	arr[startIndex], arr[randomIndex] = arr[randomIndex], arr[startIndex]
//
// 	pivot := arr[startIndex]
// 	for left < right {
// 		for left < right && arr[right] >= pivot {
// 			right--
// 		}
// 		for left < right && arr[left] <= pivot {
// 			left++
// 		}
// 		if left < right {
// 			arr[left], arr[right] = arr[right], arr[left]
// 		}
// 	}
// 	arr[left], arr[startIndex] = arr[startIndex], arr[left]
// 	if left-1 > startIndex {
// 		quickSort(arr, startIndex, left-1)
// 	}
// 	if right+1 < endIndex {
// 		quickSort(arr, right+1, endIndex)
// 	}
// }

// quickSort 填空法
func quickSort(arr []int, startIndex, endIndex int) {
	if startIndex >= endIndex {
		return
	}
	left := startIndex
	right := endIndex

	indexRange := endIndex - startIndex + 1
	randomIndex := rand.Intn(indexRange) + startIndex
	arr[startIndex], arr[randomIndex] = arr[randomIndex], arr[startIndex]

	pivot := arr[startIndex]
	for left < right {
		for left < right && arr[right] >= pivot {
			right--
		}
		if left == right {
			break
		}
		arr[left] = arr[right]
		left++
		for left < right && arr[left] <= pivot {
			left++
		}
		if left == right {
			break
		}
		arr[right] = arr[left]
		right--
	}
	arr[left] = pivot
	quickSort(arr, startIndex, left-1)
	quickSort(arr, right+1, endIndex)
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

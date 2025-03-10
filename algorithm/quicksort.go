package main

import "fmt"

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	splitData := arr[0]          // 第一个数据
	low := make([]int, 0)        // 比我小的数据
	high := make([]int, 0)       // 比我大的数据
	mid := make([]int, 0)        // 一样大的数据
	mid = append(mid, splitData) // 加入一个
	for i := 1; i < len(arr); i++ {
		if arr[i] < splitData {
			low = append(low, arr[i])
		} else if arr[i] > splitData {
			high = append(high, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}

	low, high = QuickSort(low), QuickSort(high)
	myArr := append(append(low, mid...), high...)
	return myArr
}

func main() {
	arr := []int{10, 34, 28, 48, 67, 30, 28, 65}
	fmt.Print(QuickSort(arr))
}

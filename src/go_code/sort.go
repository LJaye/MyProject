// 排序算法
// Author: 691
// Since: 2022-12-19 12:19
// File: src/go_code/sort.go
package main

import "fmt"

// 冒泡排序
func bubbleSort(arr []int) {
	// 优化：借助flag，若某一趟元素没有发生交换，则数据已经有序，不需要再继续循环判断
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 选择排序
func selectSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

// 插入排序
func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := i - 1
		tmp := arr[i]
		for j >= 0 && arr[j] > tmp {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = tmp
	}
}

// 希尔排序
func shellSort(arr []int, gap int) {
	for i := gap; i < len(arr); i++ {
		j := i - gap
		tmp := arr[i]
		for j >= 0 && arr[j] > tmp {
			arr[j+gap] = arr[j]
			j -= gap
		}
		arr[j+gap] = tmp
	}
}

// 快速排序
func quickSort(arr []int, start int, end int) {
	if start >= end {
		return
	}
	middle := partition2(arr, start, end)
	quickSort(arr, 0, middle-1)
	quickSort(arr, middle+1, end)
}

func partition(arr []int, left int, right int) int {
	i := left + 1
	j := right
	tmp := arr[left]
	for true {
		for i < right {
			if arr[i] > tmp { //找比arr[left]大的值，找不到就右移接着找
				break
			}
			i++
		}
		for j > left {
			if arr[j] < tmp { //找比arr[left]小的值，找不到就左移接着找
				break
			}
			j--
		}
		// 没有找到
		if i >= j {
			break
		}
		// 找到了，进行交换
		arr[i], arr[j] = arr[j], arr[i]
	}
	// 此时，arr[left]<arr[j]<arr[i]
	arr[left], arr[j] = arr[j], arr[left]
	return j
}

func partition2(arr []int, left int, right int) int {
	tmp := arr[left]
	for left != right {
		for left != right && arr[right] >= tmp { // 从后向前找第一个比基准小的数，加上left != right的条件防止越界
			right--
		}
		arr[left] = arr[right]
		for left != right && arr[left] <= tmp { // 从前向后找第一个比基准大的数
			left++
		}
		arr[right] = arr[left]
	}

	arr[left] = tmp
	return left // 返回左右相遇的下标
}

// 堆排序
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	arr := []int{3, 6, 2, 8, 4, 7, 888, 3, 444, 9}
	//gap := []int{5, 3, 1}
	//for _, v := range gap {
	//	shellSort(arr, v)
	//}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)

	//h := &IntHeap{3, 1, 2, 5}
	//heap.Init(h)
	//heap.Push(h, 4)
	//for h.Len() > 0 {
	//	fmt.Printf("%d ", heap.Pop(h))
	//}
}

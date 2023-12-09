package main

import "fmt"

/**
求最大子列和
*/

// 时间复杂度为O(n^3)
func maxSubSum1(arr []int, n int) int {
	maxSubSum := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			subSum := 0
			for k := i; k <= j; k++ {
				subSum += arr[k]
				if subSum > maxSubSum {
					maxSubSum = subSum
				}
			}
		}
	}
	return maxSubSum
}

// 时间复杂度为O(n^2)
func maxSubSum2(arr []int, n int) int {
	maxSubSum := 0
	for i := 0; i < n; i++ {
		subSum := 0
		for j := i; j < n; j++ {
			subSum += arr[j]
			if subSum > maxSubSum {
				maxSubSum = subSum
			}
		}
	}
	return maxSubSum
}

// 分治 时间复杂度O(  )
//func maxSubSum3(arr []int, n int) int {
//	maxSubSum := 0
//
//}

// 在线处理 时间复杂度为O(n)
func maxSubSum4(arr []int, n int) int {
	var subSum, maxSum int
	for i := 0; i < n; i++ {
		subSum += arr[i]
		if subSum < 0 { // 如果当前子列和为负
			subSum = 0 // 则不可能使后面的部分和增大，抛弃
		} else if subSum > maxSum {
			maxSum = subSum
		}
	}
	return maxSum
}

func main() {
	arr := []int{4, -3, 5, -2, -1, 2, 6, -2}
	maxSum := maxSubSum4(arr, 8)
	fmt.Println(maxSum)
}

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

// 分治 时间复杂度O(nlogn)
func maxSubarraySum(nums []int, low, high int) int {
	if low == high {
		return nums[low]
	}

	mid := (low + high) / 2

	// 分别求解左右两侧的最大子列和
	leftMax := maxSubarraySum(nums, low, mid)
	rightMax := maxSubarraySum(nums, mid+1, high)
	crossMax := maxCrossingSum(nums, low, mid, high)

	// 返回左右两侧以及跨越中点的最大子列和中的最大值
	return max(leftMax, max(rightMax, crossMax))
}

// 求解跨越中点的最大子列和
func maxCrossingSum(nums []int, low, mid, high int) int {
	leftSum := 0
	leftMaxSum := nums[mid]
	for i := mid; i >= low; i-- {
		leftSum += nums[i]
		if leftSum > leftMaxSum {
			leftMaxSum = leftSum
		}
	}

	rightSum := 0
	rightMaxSum := nums[mid+1]
	for i := mid + 1; i <= high; i++ {
		rightSum += nums[i]
		if rightSum > rightMaxSum {
			rightMaxSum = rightSum
		}
	}

	return leftMaxSum + rightMaxSum
}

// 返回两个数中的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

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
	nums := []int{4, -3, 5, -2, -1, 2, 6, -2}
	//maxSum := maxSubSum4(nums, 8)
	maxSum := maxSubarraySum(nums, 0, len(nums)-1)
	fmt.Println(maxSum)
}

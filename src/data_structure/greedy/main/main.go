package main

import (
	"fmt"
	"sort"
)

/**
贪心算法
*/

/**
1、分配问题
问题：有一群孩子和一堆饼干，每个孩子有一个饥饿度，每个饼干都有一个大小。
每个孩子只能吃最多一个饼干，且只有饼干的大小大于孩子的饥饿度时，这个孩子才能吃饱。
求解最多有多少孩子可以吃饱。

求解：因为饥饿度最小的孩子最容易吃饱，所以我们先考虑这个孩子。
为了尽量使得剩下的饼干可以满足饥饿度更大的孩子，所以我们应该把大于等于这个孩子饥饿度的、且大小最小的饼干给这个孩子。
满足了这个孩子之后，我们采取同样的策略，考虑剩下孩子里饥饿度最小的孩子，直到没有满足条件的饼干存在。
*/
func allocate(children, cookies []int) int {
	sort.Ints(children)
	sort.Ints(cookies)
	var maxChildren, j int
	for i := 0; i < len(children); i++ {
		if cookies[j] >= children[i] {
			j++
			maxChildren++
		}
	}
	return maxChildren
}

/**
2、区间问题
问题：给定多个区间，计算让这些区间互不重叠所需要移除区间的最少个数。起止相连不算重叠。

求解：在选择要保留区间时，区间的结尾十分重要: 选择的区间结尾越小，余留给其它区间的空间就越大，就越能保留更多的区间。
因此，我们采取的贪心策略为，优先保留结尾小且不相交的区间。
具体实现方法为，先把区间按照结尾的大小进行增序排序，每次选择结尾最小且和前一个选 择的区间不重叠的区间。
在样例中，排序后的数组为[[1,2],[1,3],[2,4]]。
按照我们的贪心策略，首先初始化为区间[1,2];由于[1,3] 与 [1,2] 相交，我们跳过该区间;
由于[2,4] 与 [1,2] 不相交，我们将其保留。因 此最终保留的区间为 [[1,2],[2,4]]。
*/
func interval(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	var count int
	pre := intervals[0][1] // 前一个结尾
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= pre {
			pre = intervals[i][1]
		} else {
			count++
		}
	}
	return count
}

/**
3、买卖股票的最佳时机
问题：给定一个数组，它的第 i 个元素是一支给定股票第i天的价格。
设计一个算法来计算你所能获取的最大利润。
你可以尽可能地完成更多的交易 (多次买卖一支股票)。
注意: 你不能同时参与多笔交易 (你必须在再次购买前出售掉之前的股票)。

求解：这道题的贪心思想在于如果今天买明天卖可以赚钱，那就买入。
*/
func stock(prices []int) int {
	var profit int
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			profit += prices[i+1] - prices[i]
		}
	}
	return profit
}

func main() {
	fmt.Println(allocate([]int{1, 2}, []int{1, 2, 3}))
	fmt.Println(interval([][]int{{1, 2}, {2, 4}, {1, 3}}))
	fmt.Println(stock([]int{7, 1, 5, 3, 6, 4}))
}

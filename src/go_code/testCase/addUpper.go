// 被测试的函数
// Author: 691
// Since: 2023-04-21 11:12
// File: src/go_code/testCase/addUpper.go
package testCase

func addUpper(n int) int {
	res := 0
	for i := 1; i <= n-1; i++ {
		res += i
	}
	return res
}

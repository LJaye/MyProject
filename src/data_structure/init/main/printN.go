package main

/**
打印1-n的全部正整数
*/

// 普通方式打印
func printN1(n int) {
	for i := 1; i <= n; i++ {
		println(i)
	}
	return
}

// 递归打印
func printN2(n int) {
	if n >= 1 {
		printN2(n - 1)
		println(n)
	}
	return
}

func main() {
	printN1(10)
	printN2(10)
	// 递归的代码可能很简洁，但是对空间利用效率很高，当n数字非常大的时候，递归方式可能无法运行
}

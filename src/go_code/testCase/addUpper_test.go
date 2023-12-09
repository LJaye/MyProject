// 单元测试
// Author: 691
// Since: 2023-04-21 11:21
// File: src/go_code/testCase/addUpper_test.go
package testCase

import (
	"fmt"
	"testing"
)

func TestAddUpper(t *testing.T) {
	res := addUpper(10)
	if res != 55 {
		t.Fatalf("error,expect=%v,actual=%v", 55, res)
	}
	t.Logf("right")
}

func TestHello(t *testing.T) {
	fmt.Println("hello")
}

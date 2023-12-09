// 单元测试
// Author: 691
// Since: 2023-04-21 14:26
// File: src/go_code/testCase/exercise/monster_test.go
package exercise

import (
	"testing"
)

func TestStore(t *testing.T) {
	monster := &Monster{
		Name:  "红孩儿",
		Age:   10,
		Skill: "吐火",
	}
	err := monster.Store()
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestReStore(t *testing.T) {
	var monster = &Monster{}
	err := monster.ReStore()
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf(monster.Name)
}

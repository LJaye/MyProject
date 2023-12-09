// 结构体的序列化和反序列化，保存文件
// Author: 691
// Since: 2023-04-21 14:11
// File: src/go_code/testCase/exercise/monster.go
package exercise

import (
	"encoding/json"
	"fmt"
	"os"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

// 将monster变量序列化后保存到文件中
func (m *Monster) Store() error {
	// 序列化
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("marshal error", err)
		return err
	}

	// 保存到文件
	filePath := "/Users/691/GolandProjects/monster.txt"
	err = os.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Println("write file err = ", err)
		return err
	}
	return nil
}

// 从文件读取一个序列化的monster，并反序列化
func (m *Monster) ReStore() error {
	// 读取文件
	filePath := "/Users/691/GolandProjects/monster.txt"
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("read file err = ", err)
		return err
	}

	// content []byte 反序列化
	err = json.Unmarshal(content, m)
	if err != nil {
		fmt.Println("unmarshal error", err)
		return err
	}
	return nil
}

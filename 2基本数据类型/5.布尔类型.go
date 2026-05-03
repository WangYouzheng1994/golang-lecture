package main

import "fmt"

// 默认值是false
// 不能强转其他类型
// 不能参与任何计算
func main() {
	var a bool = false
	var b bool = true

	fmt.Println(a, b)

	ab := true
	fmt.Println(ab)
}

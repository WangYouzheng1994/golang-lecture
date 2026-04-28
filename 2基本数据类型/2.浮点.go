package main

import "fmt"

// float64是默认的
// 核心原则 ：对于javaer：flat32 是float，float64是double。
// 默认值是0
func main() {
	var x float32 = 1.5
	var y float64 = 2.5

	fmt.Println(x)
	fmt.Println(y)

	// 需要强转
	y = float64(x)
	x = float32(y)

	// 模型类型是float64
	var z = 3.3
	fmt.Printf("类型：%T，值：%v", z, z)
}

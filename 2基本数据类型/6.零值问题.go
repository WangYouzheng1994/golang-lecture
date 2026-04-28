package main

import "fmt"

func main() {

	// 整型
	var a int
	var b uint64
	// 浮点
	var f32 float32
	var f64 float64
	// 布尔
	var bl bool
	// 字符
	var bt byte
	var r rune
	// 字符串
	var s string

	fmt.Printf("int: %d\n", a)
	fmt.Printf("uint64: %d\n", b)
	fmt.Printf("float32: %v\n", f32)
	fmt.Printf("float64: %v\n", f64)
	fmt.Printf("bool: %t\n", bl)
	fmt.Printf("byte: %d\n", bt)
	fmt.Printf("rune: %d\n", r)
	fmt.Printf("string: [%s]\n", s)
}

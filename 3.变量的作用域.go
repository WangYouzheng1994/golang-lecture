package main

import "fmt"

// 局部变量：定义在函数（方法）体内部；定义了必须使用
// 全局变量：定义在函数外部；定义了可以不用
func main() {
	fmt.Println(name, name2)
}

var name = "小王"
var name2 string = "xiaowang"

// 非法
//name3 := "123"

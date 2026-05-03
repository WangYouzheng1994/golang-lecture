package main

import "fmt"

// 关键点1：先进后出(栈) 延迟执行
// 关键点2：当前函数所属的作用域（调用者）结束(return)的时候才会调用defer
func main() {
	defer_m1("1")
	defer_m1("2")
	fmt.Println("main")
	defer defer_m1("3")
	defer defer_m1("4")
	defer_m2()
}

func defer_m1(str string) int {
	defer fmt.Println("defer_m1", str)
	fmt.Println("m1", str)
	return 123
}

func defer_m2() {
	defer fmt.Println("6")
	fmt.Println("10")
	defer fmt.Println("7")
}

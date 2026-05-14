package main

import "fmt"

func main() {
	fmt.Println("请输入你的年龄：")
	var age int
	fmt.Scan(&age)

	//iffun1(age)
	iffun2(age)
}

// 标准if
func iffun1(age int) {

	if age <= 0 {
		fmt.Println("未出生")
		return
	}
	if age <= 18 {
		fmt.Println("未成年")
		return
	}
	if age >= 19 && age <= 60 {
		fmt.Println("青年")
		return
	}
	fmt.Println("中年")
}

// 带初始化变量的if，作用域为if中
func iffun2(age int) {

	if newage := age; newage <= 0 {
		fmt.Println("未出生")
		return
	}
	if newage := age; newage <= 18 {
		fmt.Println("未成年")
		return
	}
	if newage := age; newage >= 19 && newage <= 60 {
		fmt.Println("青年")
		return
	}
	fmt.Println("中年")

}

package main

import "fmt"

func main() {
	one()
	fmt.Println("----------")
	two()
}

// 先声明，再赋值
func one() {
	var name string
	name = "小王"
	fmt.Println(name)
}

// 声明并赋值
func two() {
	var name string = "小王"
	fmt.Println(name)

	var name2 = "小小王"
	fmt.Println(name2)
}

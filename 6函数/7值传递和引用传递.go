package main

import "fmt"

// 关键点1：引用传递关键字*,如*int 把一个int类型的参数的指针进行传递
// 关键点2：&关键字，解析出变量的指针
func main() {
	fmt.Println("--------引用传递")
	num := 20
	fmt.Println(&num)
	fmt.Println("改以前", num)
	add_relate(&num)
	fmt.Println("改以后", num) // 成功修改 2

	fmt.Println("--------值传递")
	num2 := 20
	fmt.Println(&num2)
	fmt.Println("改以前", num2)
	add_value(num2)
	fmt.Println("改以后", num2) // 20
}

// 引用传递（指针）
func add_relate(num *int) {
	fmt.Println(num) // 内存值是一样的
	*num = 2         // 这里的修改会影响外面的num
}

// 值传递
func add_value(num int) {
	fmt.Println(&num) // 内存值是不一样的
	num = 2           // 这里的修改会影响外面的num
}

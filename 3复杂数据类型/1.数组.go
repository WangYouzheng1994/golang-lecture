package main

import "fmt"

// 值传递，值传递，和java不一样！
// 不可以扩容
func main() {
	fmt.Println("-----arr_m1()")
	arr_m1()
	fmt.Println("-----arr_m2()")
	arr_m2()
	fmt.Println("-----arr_m3()")
	arr_m3()
	fmt.Println("-----arr_m4()")
	arr_m4()
}

// 先定义再赋值
func arr_m1() {
	var a [3]int
	a = [3]int{1, 2, 3}
	fmt.Println(a)

	// 取值
	fmt.Println(a[0])
	// 根据索引赋值
	a[1] = 123
	fmt.Println(a)
}

// 短声明
func arr_m2() {
	fmt.Println("短声明")
	a := [3]int{1, 2, 3}
	fmt.Println(a)
}

// 自动推理长度
func arr_m3() {
	fmt.Println("自动推理长度")
	// 让编译器根据内容推理长度
	a := [...]int{1, 2, 3}
	fmt.Println(a)

	var b = [...]int{1, 2, 3}
	fmt.Println(b)
}

// 值传递
func arr_m4() {
	a := [...]int{1, 2, 3}
	fmt.Printf("修改前：%v\n", a)
	modify_arr(a)
	fmt.Printf("修改后：%v\n", a)
}

// 值传递
// 1. 参数必须对应要传入的长度，否则会报错
// 可以看到数组是复制了一个新的，没有影响到函数外的那个数组
func modify_arr(arr [3]int) {
	fmt.Printf("modify_arr，修改前：%v\n", arr)
	arr[0] = 3
	fmt.Printf("modify_arr，修改后：%v\n", arr)
}

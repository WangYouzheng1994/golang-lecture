package main

import (
	"errors"
	"fmt"
)

// 关键点1：多返回值是没有包裹的，和python不一样~
// 关键点2：通过_关键符号来填充
func main() {
	fmt.Println("fun2()---------")
	fmt.Printf("返回值：%v, 类型：%T", fun2(), fun2())
	fmt.Println("fun3()---------")
	fmt.Println(fun3())
	// 短定义接受两个参数
	i, err := fun3()
	// 接收两个参数
	var i1, err1 = fun3()
	// 只接收i2这个参数
	var i2, _ = fun3()
	fmt.Println(i2)
	fmt.Println(i, err)
	fmt.Println(i1, err1)
	fmt.Println("fun4()---------")
	fmt.Printf("返回值：%v, 类型：%T", fun4(), fun4())
}

// 无返回值
func fun1() {
	return // 也可以不写
}

// 单返回值
func fun2() int {
	return 1
}

// 多返回值
func fun3() (int, error) {
	return 0, errors.New("错误")
}

// 命名返回值
func fun4() (res string) {
	return // 相当于先定义再赋值
	//return "abc"
}

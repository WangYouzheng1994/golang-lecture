package main

import "fmt"

// 关键点1：别名，不可以绑定方法
// 关键点2： 可以和原始类型直接比较
func main() {
	// 类型别名，打印它的类型还是原始类型
	fmt.Printf("%T %T \n", SuccessCode, SuccessAliasCode) // main.MyCode int
	// 可以直接和原始类型比较
	var i int
	fmt.Println(SuccessAliasCode == i)
	fmt.Println(int(SuccessCode) == i) // 必须转换之后才能和原始类型比较
}

type AliasCode = int
type MyCode int

const (
	SuccessCode      MyCode    = 0
	SuccessAliasCode AliasCode = 0
)

// MyCodeMethod 自定义类型可以绑定自定义方法
func (m MyCode) MyCodeMethod() {

}

// MyAliasCodeMethod 类型别名 不可以绑定方法
//func (m AliasCode) MyAliasCodeMethod() {
//
//}

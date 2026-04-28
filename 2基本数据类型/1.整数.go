package main

import "fmt"

// 没有特殊要求的时候直接用 int，系统会根据int32、int64进行自适应
// 核心原则1 ：uint 是无符号，int是有符号
// 核心原则2 ：对于javaer：int32 是int，int64是long。
// PS: 用过mongo的对这一套会比较熟悉，主要是内存敏感、以及bit定义的时候可以用这种细分数据类型
func main() {
	int_m1()
}
func int_m1() {
	//  无符号整数 - start
	var n1 uint8 = 2
	var n2 uint16 = 2
	var n3 uint32 = 2
	var n4 uint64 = 2
	var n5 uint = 2
	//  无符号整数 - end

	// 有符号整数 - start
	//var n6 int8 = 2
	//var n7 int16 = 2
	//var n8 int32 = 2
	//var n9 int64 = 2
	//var n10 int = 2
	// 有符号整数 - end
	fmt.Println(n1)
	fmt.Println("-------")
	fmt.Println(n2)
	fmt.Println("-------")
	fmt.Println(n3)
	fmt.Println("-------")
	fmt.Println(n4)
	fmt.Println("-------")
	fmt.Println(n5)
	fmt.Println("-------")

	var n6 int8 = 2
	fmt.Println(n6)
	fmt.Println("-------")
	// 不同类型不能自动强转，必须手动。
	n1 = uint8(n6)
	fmt.Println(n6)

	// 自适应，默认就是int
	var n11 = 2
	fmt.Printf("类型: %T，值：%v", n11, n11)

}

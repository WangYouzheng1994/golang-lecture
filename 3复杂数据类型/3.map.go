package main

import "fmt"

// 键值对的map
// 关键点1：key和value在声明的时候要求固定下，不可以像java在整个容器范围的key、value可以是不同的类型
// 关键点2：引用传递
// 关键点3：栈中存的是指向堆的开头，底层是一个连续的内存，通过哈希表数组维护了数据key。因此key、value都在堆。99%等价于java map
func main() {
	fmt.Println("--------map_m1----------")
	map_m1()
	fmt.Println("--------map_m2----------")
	map_m2()
}

func map_m1() {
	// 声明
	var m1 map[string]string
	// 初始化1
	m1 = make(map[string]string)
	fmt.Println(m1)
	// 设置值
	m1["name"] = "枫枫"
	fmt.Println(m1)
}

func map_m2() {
	// 初始化2
	var m1 map[string]string = map[string]string{}
	// 设置值
	m1["name"] = "枫枫"
	fmt.Println(m1)
	fmt.Printf("取一个不存在的key,类型是：%T, 值是：%v", m1["abc"], m1["abc"])
}

package main

import "fmt"

// 值传递
func (s Student) SetAge(age int) {
	s.Age = age
}

// 指针（引用）传递
func (s *Student) SetAge1(age int) {
	s.Age = age
}

// 本案例展示结构体函数的值传递和指针传递的区别
func main() {
	s := Student{
		Name: "枫枫",
		Age:  21,
	}
	// 值传递 无法修改s
	s.SetAge(18)
	fmt.Println(s.Age) // 21
	// 引用传递 修改s
	s.SetAge1(18)
	fmt.Println(s.Age) // 18
}

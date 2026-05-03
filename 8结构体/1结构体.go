package main

import "fmt"

// 成员变量，大写 = 包外可访问（public），小写 = 仅本包私有（private）

func main() {
	s := Student{
		Name: "小王",
		Age:  1,
	}
	s.Name = "小小王" // 修改值
	s.PrintInfo()
}

// Student 定义结构体Student
type Student struct {
	Name string
	Age  int
}

// Student2 定义结构体
type Student2 struct {
	Name string
	Age  int
}

// PrintInfo 给结构体Student绑定一个方法
func (s Student) PrintInfo() {
	fmt.Printf("name:%s age:%d\n", s.Name, s.Age)
}

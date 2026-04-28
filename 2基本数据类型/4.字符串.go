package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 字符串用双引号定义
// 多行字符串用“定义，可以带着编辑器中的缩进进行原样输出
// 默认值是空串
func main() {
	//str_m1()
	str_m2()
}

// 定义
func str_m1() {
	// 单行字符串
	var str = "小王"

	// 多行字符串
	var multi = `
	小
	小
	王
	`
	fmt.Println(str)
	fmt.Println(multi)
}

func str_m2() {
	var str = "小王"
	// 长度
	fmt.Println(len(str))
	// 拼接
	fmt.Println(str + "长度：" + strconv.Itoa(len(str)))
	// 分割
	str = "小,王"
	fmt.Println(strings.Split(str, ","))
	// 包含
	fmt.Println(strings.Contains(str, ","))
}

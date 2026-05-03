package main

import "fmt"

// 关键点1：swtich 自带break语义
// 关键点2：swtich fallthrough 穿透
func main() {
	fmt.Println("请输入你的年龄：")
	var age int
	fmt.Scan(&age)
	fmt.Println("----------switch_m1----------")
	switch_m1(age)
	fmt.Println("----------switch_m2----------")
	switch_m2(age)
}

func switch_m1(age int) {

	switch {
	case age <= 0:
		fmt.Println("未出生")
	case age <= 18:
		fmt.Println("未成年")
	case age <= 35:
		fmt.Println("青年")
	default:
		fmt.Println("中年")
	}
}

func switch_m2(age int) {
	switch age {
	case 18:
		fmt.Println("年轻")
	case 45:
		fmt.Println("中年")
		// 穿透
		fallthrough
	default:
		fmt.Println("退休")
	}
}

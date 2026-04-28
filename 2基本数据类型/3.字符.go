package main

import "fmt"

// 字符：是定长的，单引号定义。rune是unicode字符（中文、emoji）；byte是UTF8字节（数字、字母）
// 关键点1：默认的是rune；  rune(int32) var a,b = 'a', '中';
// 关键点2：需要显式指定为byte； byte(uint8) var a,b byte = 'a', '32'
// 默认值是0
func main() {
	rune_m1()
	fmt.Println("---------------")
	byte_m1()
}

// 通过打印可以看出底层是数字（unicode字符集位置）
func rune_m1() {
	var c1 = 'a'
	var c2 = 97
	fmt.Println(c1) // 直接打印都是数字
	fmt.Println(c2)

	fmt.Printf("%c %c\n", c1, c2) // 以字符的格式打印

	var r1 rune = '中'
	fmt.Printf("%c\n", r1)

	fmt.Printf("类型是：%T\n %T\n %T\n", c1, c2, r1)
}

func byte_m1() {
	// byte == uint8
	var a, b byte = 'a', '3'
	fmt.Println(a, b)
	fmt.Printf("类型是：%T\n %T", a, b)
}

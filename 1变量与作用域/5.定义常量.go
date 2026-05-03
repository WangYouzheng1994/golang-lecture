package main

import "fmt"

const name3 = "abcd"

const (
	name5 = "5555555"
	name6 = "6666666"
)

func main() {
	const name string = "ab"
	const name2 = "abc"

	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(name3)
	fmt.Println(name4)
	fmt.Println(name5)
	fmt.Println(name6)
}

const name4 = "abcde"

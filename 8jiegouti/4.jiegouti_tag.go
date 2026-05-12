package main

import (
	"encoding/json"
	"fmt"
)

// 关键点1：如果成员变量名字小写，那么无法序列化
// 关键点2：如果不序列化，可以指定为横杠-
type Student struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	height int    `json:"heig"` // 小写无法序列化
	Weight int    `json:wei`
	//Password string `json:"-"` // 不序列化
}

func main() {
	s := Student{
		Name: "枫枫",
		Age:  21,
		//Password: "Password",
		height: 155,
		Weight: 156,
	}

	// 序列化

	byteData, _ := json.Marshal(s)
	fmt.Println(string(byteData))
}

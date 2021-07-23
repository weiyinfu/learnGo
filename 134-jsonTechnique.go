package main

import (
	"fmt"
)

type User struct {
	//使用json注解
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type User2 struct {
	//数字转字符串
	ID int64 `json:"id,string"`
}
type User3 struct {
	//忽略空值
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
type User4 struct {
	//忽略key
	ID   int    `json:"id"`
	Name string `json:"-"`
}

func main() {
	fmt.Println("haha")
}

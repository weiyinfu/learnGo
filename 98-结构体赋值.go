package main

import "fmt"

func main() {
	type haha struct {
		name string
	}
	//结构体赋值是深拷贝
	x := haha{name: "one"}
	y := x
	y.name = "two"
	fmt.Println(x.name)//输出one，说明结构体赋值是深度赋值
}

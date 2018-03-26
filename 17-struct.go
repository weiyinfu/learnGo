package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"weidiao", 25})
	fmt.Println(person{name: "weidiao", age: 25})
	fmt.Println(person{name: "weidiao"})
	one := person{"weidiao", 25}
	two := &one
	(*two).age = 26 //*其实是可以省略的
	fmt.Println(one)
	two.age = 27
	fmt.Println(one)

	//结构体传值和传址要区分清楚
	func(x person) {
		x.name = "haha"
	}(one)
	fmt.Println(one)
	func(x *person) {
		x.name = "haha"
	}(&one)
	fmt.Println(one)
}

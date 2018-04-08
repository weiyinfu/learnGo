package main

import "fmt"

type Node86 struct {
	age int
}

//结构体传址
func (node *Node86) plus(x int) {
	node.age += x
}

//结构体传值
func (node Node86) plus2(x int) {
	node.age += x
}
func main() {
	a := Node86{2}
	fmt.Println(a)
	a.plus(1)
	fmt.Println(a)
	a.plus2(1)
	fmt.Println(a)
}

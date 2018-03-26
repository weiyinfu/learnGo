package main

import "fmt"

type Animal interface {
	shout()
}
type Dog struct {
}

func (a Dog) shout() {
	fmt.Println("wang wang wang")
}

type Cat struct {
}

func (a Cat) shout() {
	fmt.Println("miao miao miao")
}
func main()  {
	d:=Dog{}
	c:=Cat{}
	c.shout()
	d.shout()
}

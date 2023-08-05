package main

import "fmt"

type Animal19 interface {
	shout()
}
type Dog struct {
}

func (a Dog) shout() {
	fmt.Println("wang wang wang")
}

type Cat19 struct {
}

func (a Cat19) shout() {
	fmt.Println("miao miao miao")
}
func main() {
	var d Animal19 = Dog{}
	var c Animal19 = Cat19{}
	c.shout()
	d.shout()
	d = c
	d.shout()
}

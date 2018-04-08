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
	var d Animal =Dog{}
	var c Animal=Cat{}
	c.shout()
	d.shout()
	d=c
	d.shout()
}

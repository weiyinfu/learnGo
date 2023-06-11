package main

import "fmt"

type One struct {
	x int
}

func (x One) handle() {
	//x.x++
	fmt.Println(&x.x, "value=", x.x)
}

type O interface {
	handle()
}

func main() {
	x := One{}
	x.handle()
	var xx O
	xx = x
	fmt.Println(&xx)
	xx.handle()
	x.handle()
}

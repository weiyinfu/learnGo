package main

import "fmt"

/**
虽然go语言如此灵活，以至于它并不迫切需要构造函数，但是go还是提供了new关键字
*/
type Dog2 struct {
	name string
}

func newDog() Dog2 {
	dog := new(Dog2) //new出来的是一个指针
	dog.name = "weidiao"
	return *dog
}
func main() {
	fmt.Println(newDog())
}

package main

import "fmt"

type Animal3 struct {
}

func (ani3 Animal3) shout() {
	fmt.Println("shout")
}

type Person struct {
}

func (person Person) run() {
	fmt.Println("run")
}

type Man struct {
	Animal3
	Person
}

func main() {
	man := Man{}
	man.shout()
	man.run()
}

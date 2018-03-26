package main

import "fmt"

type Animal2 struct {
	name string
}

func (ani Animal2)shout()  {
	fmt.Println(ani)
}

type Bird struct {
	Animal2
}

func main() {
	Bird{Animal2{"bird"}}.shout()
}

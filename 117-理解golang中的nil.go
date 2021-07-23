package main

import "fmt"

type C struct {
	name string
}

func main() {
	var a *C
	fmt.Println(a == nil, a == (*C)(nil), a)
}

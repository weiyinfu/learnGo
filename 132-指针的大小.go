package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x *int
	fmt.Println(unsafe.Sizeof(x))
	s := "weiyinfu"
	fmt.Println(len(s), unsafe.Sizeof(s))
}

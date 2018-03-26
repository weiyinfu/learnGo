package main

import "fmt"

func main() {
	s := 0
	x := 0
forLoop:
	if x < 10 {
		s += x
		x++
		goto forLoop

	}
	fmt.Println(s)
}

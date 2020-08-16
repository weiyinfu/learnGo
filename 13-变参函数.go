package main

import "fmt"

func Sum(a ...int) int {
	s := 0
	for _, i := range a {
		s += i
	}
	return s
}
func main() {
	fmt.Println(Sum(1, 2, 3, 4))
}

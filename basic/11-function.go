package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

//命名返回值
func sum(x, y, z int) (s, prod int) {
	s = x + y + z
	prod = x * y * z
	return
}
func main() {
	fmt.Println(add(3, 4))
	fmt.Println(sum(1, 2, 3))
}

package main

import "fmt"

func main() {
	fmt.Println(func(x, y int, f func(int, int) int) int {
		return x + y + f(x, y)
	}(1, 2, func(x, y int) int { return x - y }))
	for i := 0; i < 10; i++ {
		fmt.Print(func(x int) int { return x + i }(i))
	}
}

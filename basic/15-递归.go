package main

import "fmt"

func fac(x int) int {
	if x == 0 || x == 1 {
		return 1
	} else {
		return x * fac(x-1)
	}
}
func main() {
	fmt.Println(fac(5))
}

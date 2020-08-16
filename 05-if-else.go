package main

import "math/rand"

func main() {
	if x := rand.Int(); x%2 == 0 {
		print("x是个偶数")
	} else {
		print("x是个奇数")
	}
}

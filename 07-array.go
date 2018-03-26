package main

import "fmt"

func main() {
	var a [5]int
	fmt.Print(a)
	a[0] = 3
	fmt.Println(a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	var c [2][3]int
	k := 0
	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c[i]); j++ {
			c[i][j] = k
			k++
		}
	}
	fmt.Println(c)
}

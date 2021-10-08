package main

import "fmt"

func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}
func main() {
	a := []int{1, 2}
	swap(&a[0], &a[1])
	fmt.Println(a)
	a[0], a[1] = a[1], a[0]
	fmt.Println(a)
}

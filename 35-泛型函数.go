package main

import (
	"fmt"
)

/**
Go本身是不支持泛型的
*/
type Comparable interface {
	Len() int
	Compare(int, int) int
	Swap(int, int)
}

type IntArray []int

func (ar IntArray) Len() int {
	return len(ar)
}
func (ar IntArray) Compare(x, y int) int {
	return ar[x] - ar[y]
}
func (ar IntArray) Swap(x, y int) {
	ar[x], ar[y] = ar[y], ar[x]
}
func main() {
	swapSort := func(a Comparable) {
		for i := 0; i < a.Len(); i++ {
			for j := i; j < a.Len(); j++ {
				if a.Compare(i, j) > 0 {
					a.Swap(i, j)
				}
			}
		}
	}

	a := IntArray{1, 5, 3, 2, 4}
	swapSort(a)
	for _, x := range a {
		fmt.Print(x, " ")
	}
}

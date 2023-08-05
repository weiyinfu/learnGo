package main

import "fmt"

func main() {
	a := []int{1, 3, 2, 4, 6}
	quickSort(a)
	fmt.Println(a)
}
func handle(a []int, beg, end int) {
	if beg >= end {
		return
	}
	i := beg
	j := end
	v := a[beg]
	for i < j {
		for i < j && a[j] > v {
			j--
		}
		if i >= j {
			break
		}
		a[i] = a[j]
		for i < j && a[i] < v {
			i++
		}
		if i >= j {
			break
		}
		a[j] = a[i]
	}
	a[j] = v
	handle(a, beg, i-1)
	handle(a, i+1, end)
}
func quickSort(a []int) {
	handle(a, 0, len(a)-1)
}

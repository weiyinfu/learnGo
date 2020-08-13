package main

import "fmt"
/**
slice追加元素的时候，可以自动扩容
扩容之后，切片不会触发复制操作
*/
func appendChange(a []int) {
	a = append(a, 3)
	a = append(a, 4)
	a[0] = 999
	fmt.Println("%p:%v", a, a)
}
func main() {
	a := make([]int, 0, 2)
	a = append(a, 1)
	a = append(a, 2)
	appendChange(a)
	fmt.Println("%p:%v", a, a)
}

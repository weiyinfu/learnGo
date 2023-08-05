package main

import "fmt"

//you don't need class at all
type Rectangle struct {
	width, height int
}

func (r *Rectangle) area() int {
	ans := r.width * r.height
	return ans
}
func (r Rectangle) length() int {
	return (r.width + r.height) << 1
}
func main() {
	r := Rectangle{3, 4}
	fmt.Println(r.area(), r.length())
	fmt.Println(r)
	rr := &r
	fmt.Println(rr.area(), rr.length())
}

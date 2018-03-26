package main

import "fmt"

//you don't need class at all
type Rectangle struct {
	width, height int
}

func (r *Rectangle) area() int {
	return r.width * r.height
}
func (r Rectangle) length() int {
	return (r.width + r.height) << 1
}
func main() {
	r := Rectangle{3, 4}
	fmt.Println(r.area(), r.length())
	rr := &r
	fmt.Println(rr.area(), rr.length())
}

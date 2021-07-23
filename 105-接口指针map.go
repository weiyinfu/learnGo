package main

import "fmt"

type haha105 interface {
	haha()
}
type one struct {
	id int
}

func (x *one) haha() {

}
func main() {
	a := map[haha105]bool{}
	x := &one{id: 3}
	y := one{id: 3}
	a[x] = true
	a[&y] = true
	fmt.Println(a)
	/*
		结构体map
	*/
	b := map[one]bool{}
	b[*x] = true
	b[y] = true
	fmt.Println(b)
}

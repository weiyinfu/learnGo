package main

import "fmt"

type interface104 interface {
	haha()
}
type Dog104 struct {
	id int
}

func (x Dog104) haha() {
	fmt.Println(x.id)
}
func get() []interface104 {
	var a []interface104
	for i := 0; i < 10; i++ {
		a = append(a, &Dog104{})
	}
	return a
}
func main() {
	x := get()
	//y := x[0].(Dog104)
	//以上代码会报错
	y := x[0].(*Dog104)
	y.id = 100
	x[0].haha()
}

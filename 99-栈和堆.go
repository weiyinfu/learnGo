package main

import "fmt"
/**
在C++中，函数返回局部变量的指针，这是典型的错误。
在Go中，编译器会进行一种名为"逃逸分析"的处理，判断局部变量的生命周期是否超出了函数的局部作用域。
如果超出了，则将结构体对象在堆上分配。
否则，将结构体对象在栈上分配。
*/
type Apple struct {
	age int
}

func haha10() *Apple {
	var x int //8B空间
	fmt.Println("x的地址", &x)
	var w int
	fmt.Println("w的地址", &w)
	/**
	a是分配在堆上的还是分配在栈上的？
	a占了32B空间
	*/
	a := Apple{
		age: 10,
	}
	var y int
	fmt.Println("y的地址", &y)
	var z int
	fmt.Println("z的地址", &z)
	b := &a
	fmt.Println("b的地址：返回之前的地址", &b)
	var m int
	fmt.Println("m的地址", &m)
	return &a
}
func main() {
	a := haha10()
	fmt.Println("返回值后的地址",&a)
	println(a.age)
}

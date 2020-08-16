package main

import (
	"fmt"
	"math"
)

const s string = "haha"

func main() {
	fmt.Println(s)
	const n = 1000
	fmt.Println(n)
	const m = 2e30 / n
	fmt.Println(m, math.Sin(m))

	const (
		SUN  = 1
		MOON = 2
	)
	sun := 1
	switch sun {
	case SUN:
		fmt.Println("sun")
	case MOON:
		fmt.Println("moon")
	}
	//iota，特殊常量，可以认为是一个可以被编译器修改的常量。
	//iota就是为枚举而生的
	const (
		one = iota
		two = iota
	)
	fmt.Println(one, two)
	const (
		three = iota
		four  = iota
	)
	fmt.Println(three, four)

	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)

	fmt.Println(a, b, c, d, e, f, g, h, i)
	const (
		ii = 1 << iota
		j  = 3 << iota
		k  //3<<2
		l  //3<<3
	)
	fmt.Println(ii, j, k, l)

}

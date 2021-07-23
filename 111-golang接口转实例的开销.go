package main

import (
	"fmt"
	"time"
)

/**
测试golang接口类型转换成实例类型的开销
转换的开销几乎可以忽略不计
*/
type Animal interface {
	shout()
}
type Cat struct {
	id int
}

func (self *Cat) shout() {

}
func getCats(n int) []Animal {
	var a []Animal
	for i := 0; i < n; i++ {
		a = append(a, &Cat{id: 0})
	}
	return a
}
func setCats(a []Animal) {
	for _, i := range a {
		ii := i.(*Cat)
		ii.shout()
		ii.id = 0
	}
}
func setCats2(a []*Cat) {
	for _, ii := range a {
		ii.shout()
		ii.id = 0
	}
}
func timeit(f func(), t func(), times int) {
	a := []func(){f, t}
	used := []time.Duration{0, 0}
	for i := 0; i < times; i++ {
		beginTime := time.Now()
		a[i%2]()
		endTime := time.Now()
		u := endTime.Sub(beginTime)
		used[i%2] += u
	}
	fmt.Println(used)
}
func main() {
	a := getCats(10000)
	var b []*Cat
	for _, i := range a {
		b = append(b, i.(*Cat))
	}
	n := 5000000
	timeit(func() {
		setCats(a)
	}, func() {
		setCats2(b)
	}, n)
}

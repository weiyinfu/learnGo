package main

import (
	"sync"
	"time"
	"fmt"
)

var a string
var once sync.Once

func setup() {
	fmt.Println("setup")
	a = "hello, world"
}

func doprint() {
	//多次调用此函数，setup只被调用一次
	once.Do(setup)
	fmt.Println(a)
}

func main() {
	go doprint()
	go doprint()
	time.Sleep(3 * time.Second)
}

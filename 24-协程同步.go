package main

import (
	"fmt"
	"time"
)

func main() {
	worker := func(x chan bool) {
		fmt.Println(x)
		fmt.Println("I will sleep")
		time.Sleep(time.Second)
		fmt.Println(time.Nanosecond)

		x <- true
	}
	done := make(chan bool)
	go worker(done)
	<-done //从done中读取数据，直到协程执行完毕
}

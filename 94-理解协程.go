package main

import (
	"fmt"
	"runtime"
	"time"
)

/**
Java中的单线程是顺序执行的
单线程内的协程是交替执行的
*/
func main() {
	runtime.GOMAXPROCS(1)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("task1", time.Now())
			time.Sleep(time.Second)
		}
	}()
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("task2", time.Now())
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("task3", time.Now())
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Second * 15)
}

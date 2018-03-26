package main

import (
	"math/rand"
	"fmt"
	"time"
)

/**
真是一件神奇的事情,下面的代码哪里不对呢?
*/
func main() {
	type Task struct {
		x, y int
		sum  chan int
	}
	jobs := make(chan Task, 100)
	tasks := [10]Task{}
	for i := 0; i < 103; i++ {
		go func(jobs <-chan Task) {
			for j := range jobs {
				fmt.Println("haha")
				j.sum <- j.x + j.y
			}
		}(jobs)
	}

	for j := 0; j < 10; j++ {
		task := Task{x: rand.Intn(10), y: rand.Intn(10)}
		jobs <- task
		tasks[j] = task
	}
	close(jobs)
	time.Sleep(time.Second * 5)
	for _, i := range tasks {
		fmt.Printf("%d+%d=%d\n", i.x, i.y, <-i.sum)
	}

}

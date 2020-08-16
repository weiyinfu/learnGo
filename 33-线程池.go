package main

import (
	"fmt"
	"math/rand"
	"sync"
)

//找bug：如何解决本程序中的协程死锁
func main() {
	type Task struct {
		x, y int
		sum  chan int
	}
	jobs := make(chan *Task, 100)
	tasks := [10]*Task{}
	consumerCount := 2
	producerCount := 3
	wg:=sync.WaitGroup{}
	//消费者
	for i := 0; i < consumerCount; i++ {
		wg.Add(1)
		go func(jobs <-chan *Task) {
			for {
				j, more := <-jobs
				if !more {
					break
				}
				fmt.Println("Got Task")
				j.sum <- j.x + j.y
				fmt.Println("write over")
			}
			fmt.Println("工作干完了")
			wg.Done()
		}(jobs)
	}
	//生产者
	for j := 0; j < producerCount; j++ {
		task := Task{x: rand.Intn(10), y: rand.Intn(10), sum: make(chan int)}
		jobs <- &task
		tasks[j] = &task
	}
	close(jobs)
	fmt.Println("closed jobs")
	wg.Wait()
	for _, i := range tasks {
		if i == nil {
			continue
		}
		fmt.Printf("%d+%d=%d\n", i.x, i.y, <-i.sum)
	}
}

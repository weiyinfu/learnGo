package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type task struct {
	x, y int
	z    int
}

func main() {
	buf := make(chan task)
	exit := make(chan struct{})
	var wg sync.WaitGroup//waitGroup用于等待若干个进程结束，每次可以Add，也可以Done
	//producer
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			for {
				select {
				case <-exit:
					wg.Done()
					return
				default:
					buf <- task{rand.Intn(10), rand.Intn(10), 0}
				}
			}
		}()
	}
	//consumer
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for {
				select {
				case <-exit:
					wg.Done()
					return
				default:
					it := <-buf
					fmt.Println(it.x, it.y, it.x+it.y)
				}
			}
		}()
	}
	time.Sleep(time.Minute)
	close(exit)
	wg.Wait()
}

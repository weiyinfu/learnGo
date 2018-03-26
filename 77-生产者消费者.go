package main

import (
	"sync"
	"math/rand"
	"fmt"
	"time"
)

type data struct {
	x, y int
	z    int
}

func main() {
	buf := make(chan data)
	exit := make(chan struct{})
	var wg sync.WaitGroup
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
					buf <- data{rand.Intn(10), rand.Intn(10), 0}
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

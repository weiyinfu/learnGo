package main

import (
	"fmt"
	"go.uber.org/atomic"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestOne(t *testing.T) {
	x := 0
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x++
			}
		}()
	}
	time.Sleep(time.Second * 2)
	fmt.Println(x)
}

func Test2(t *testing.T) {
	x := 0
	runtime.GOMAXPROCS(1)
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x++
			}
		}()
	}
	time.Sleep(time.Second * 2)
	fmt.Println(x)
}
func Test3(t *testing.T) {
	l := sync.Mutex{}
	x := 0
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				l.Lock()
				x++
				l.Unlock()
			}
		}()
	}
	time.Sleep(time.Second * 2)
	fmt.Println(x)
}
func Test4(t *testing.T) {
	x := atomic.Int32{}
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x.Add(1)
			}
		}()
	}
	time.Sleep(time.Second * 2)
	fmt.Println(x.Load())
}
func Test5(t *testing.T) {
	x := 0
	send := make(chan int)
	done := make(chan bool)
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				send <- 1
			}
		}()
	}
	go func() {
		for {
			select {
			case <-done:
				{
					break
				}
			case <-send:
				{
					x++
				}
			}
		}
	}()
	time.Sleep(time.Second * 2)
	done <- true
	fmt.Println(x)
}

func Test6(t *testing.T) {

}

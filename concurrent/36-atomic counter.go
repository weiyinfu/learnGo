package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var ops uint64
	for i := 0; i < 50; i++ {
		go func() {
			atomic.AddUint64(&ops, 1)
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(ops)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println(opsFinal, ops)
}

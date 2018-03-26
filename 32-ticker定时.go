package main

import (
	"time"
	"fmt"
)

func main() {
	t := time.NewTicker(time.Second)
	go func() {
		for i := range t.C {
			fmt.Println("tick at", i)
		}
	}()
	time.Sleep(5 * time.Second)
	t.Stop()
}

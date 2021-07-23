package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(time.Second)
	go func() {
		for i := range t.C {
			fmt.Println("tick at", i.String(), time.Now().String())
		}
	}()
	time.Sleep(5 * time.Second)
	t.Stop()
}

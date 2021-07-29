package main

import (
	"fmt"
	"time"
)

func main() {
	x := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		x <- "haha"
	}()
	select {
	case ans := <-x:
		fmt.Println(ans)
	case <-time.After(2 * time.Second):
		{
			fmt.Println("在规定时间内未能完成任务")
		}
	}
}

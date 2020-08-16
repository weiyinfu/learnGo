package main

import (
	"fmt"
	"time"
)

/*
两个协程交替打印ping、pong，直到接收到用户输入
*/
func main() {
	ping := make(chan string, 1)
	pong := make(chan string, 1)
	go func(ping chan<- string, pong <-chan string) {
		//此函数有两个管道参数，第一个ping是一个写管道，第二个参数pong是一个读管道
		for {
			ping <- "ping"
			fmt.Println(<-pong)
			time.Sleep(time.Second * 2)
		}
	}(ping, pong)
	go func(ping <-chan string, pong chan<- string) {
		for {
			pong <- "pong"
			fmt.Println(<-ping)
			time.Sleep(time.Second * 3)
		}
	}(ping, pong)
	_, _ = fmt.Scanln()
}

package main

import (
	"fmt"
	"time"
)

func main() {
	ping := make(chan string, 1)
	pong := make(chan string, 1)
	go func(ping chan<- string, pong <-chan string) {
		for {
			ping <- "ping"
			fmt.Println(<-pong)
			time.Sleep(time.Second*2)
		}
	}(ping, pong)
	go func(ping <-chan string, pong chan<- string) {
		for {
			pong <- "pong"
			fmt.Println(<-ping)
			time.Sleep(time.Second*5)
		}
	}(ping, pong)
	fmt.Scanln()
}

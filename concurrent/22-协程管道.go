package main

import "fmt"

func main() {
	message := make(chan string)
	go func() { message <- "ping" }()
	msg := <-message //从管道中读取信息
	fmt.Println(msg)
}

package main

import (
	"fmt"
	"time"
)

/**
* 使用不阻塞通道实现
*/
func main() {
	message := make(chan string)
	go func() {
		message <- "haha"
	}()
	//如果主线程睡一觉,下面就会打印haha,否则就会直接跳过
	time.Sleep(time.Second)
	select {
	case msg := <-message:
		{
			fmt.Println(msg)
		}
	default:
		fmt.Println("直接跳过")

	}
}

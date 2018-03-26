package main

import (
	"time"
	"fmt"
	"math/rand"
)

/**
老师出题,同学们抢答
可是为啥同学们回答速度越来越快呢?
**/
func main() {
	one := make(chan string)
	two := make(chan string)
	for {
		go func() {
			t := time.Second * time.Duration(rand.Uint32()%5+2)
			fmt.Println("first run", t)
			time.Sleep(t)
			one <- "one"
		}()
		go func() {
			t := time.Second * time.Duration(rand.Uint32()%5+2)
			fmt.Println("second run", t)
			time.Sleep(t)
			two <- "two"
		}()
		select {
		case msg1 := <-one:
			{
				fmt.Println("first is right", msg1)
				one <- nil
			}
		case <-two:
			fmt.Println("second is right")
			two <- nil
		}
		fmt.Println("我要出题了")
	}
}

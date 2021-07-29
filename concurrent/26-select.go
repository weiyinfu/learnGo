package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
老师出题,同学们抢答
可是为啥同学们回答速度越来越快呢?

使用select来选择最快返回的那一条
**/
func main() {
	one := make(chan string)
	two := make(chan string)
	done := make(chan bool)
	go func() {
		t := time.Second * time.Duration(rand.Uint32()%5+2)
		fmt.Println("first run", t)
		time.Sleep(t)
		select {
		case <-done:
			{
				//比赛已经结束了
			}
		default:
			one <- "one"
		}
	}()
	go func() {
		t := time.Second * time.Duration(rand.Uint32()%5+2)
		fmt.Println("second run", t)
		time.Sleep(t)
		select {
		case <-done:
			{
				//比赛已经结束了
			}
		default:
			two <- "two"
		}
	}()
	select {
	case msg1 := <-one:
		{
			done <- true
			fmt.Println("first is right", msg1)
		}
	case <-two:
		done <- true
		fmt.Println("second is right")
	}
}

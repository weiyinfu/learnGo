package main

import (
	"fmt"
	"time"
)

var chA = make(chan int)
var chB = make(chan int)
var over = make(chan int)

func printA() {
	for {
		select {
		case <-chA:
			{
				fmt.Println("A")
				chB <- 1
				time.Sleep(time.Second * 2)
			}
		default:
			{

			}
		}
	}

}
func printB() {
	for {
		select {
		case <-chB:
			{
				fmt.Println("B")
				chA <- 1
				time.Sleep(time.Second * 2)
			}
		default:
			{

			}
		}
	}

}
func main() {
	go func() {
		printA()
	}()
	go func() {
		printB()
	}()
	chA <- 1
	<-over
}

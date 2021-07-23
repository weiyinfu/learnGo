package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		//没有这句话，就会直接panic
		//defer func() {
		//	recover()
		//}()
		fmt.Println("hahaha")
		panic("baga")
	}()
	fmt.Println("over")
	time.Sleep(time.Second * 10)
}

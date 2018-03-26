package main

import (
	"time"
	"fmt"
)

func main() {
	t := time.NewTimer(time.Second * 3)
	<-t.C
	fmt.Println("我刚才睡了一觉")

	tt := time.NewTimer(time.Second * 10)
	go func() {
		<-tt.C
		fmt.Println("子线程睡醒了")
	}()
	stop := tt.Stop()
	if stop {
		fmt.Println("子线程被叫醒了")
	} else {
		fmt.Println("子线程执行完了")
	}
}

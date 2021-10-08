package main

import (
	"fmt"
	"testing"
	"time"
)

/**
注意：如果执行One、Two、Three，因为它们并没有占用CPU，所以CPU占用时长为0.因此cpu profile里面没有它们。
*/
func One() {
	time.Sleep(time.Second)
}
func Two() {
	time.Sleep(time.Second * 2)
}
func Three() {
	time.Sleep(time.Second * 3)
}

func Test_Main(t *testing.T) {
	for i := 0; i < 2; i++ {
		fmt.Println("running test")
		One()
		Two()
		Three()
		Four()
		Five()
	}
}

package main

import (
	"fmt"
	"time"
	"unsafe"
)

func get(s string) {
	fmt.Println("[get]", s)
	time.Sleep(time.Second * 4)
	fmt.Println("[get]after", s)
}
func main() {
	s := []byte{65, 66, 67, 68}
	p := unsafe.Pointer(&s)
	ss := (*string)(p)
	region := *ss
	fmt.Println("[main]", *ss)
	go func() {
		get(region)
	}()
	time.Sleep(time.Second)
	s[1] = 70
	fmt.Println("[main]after", *ss)
	time.Sleep(time.Second * 5)
}

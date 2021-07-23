package main

import (
	"fmt"
)

func haha() string {
	fmt.Println("haha is called")
	return "haha"
}
func main() {
	//没有这句话，就会直接panic
	//defer中如果defer的是一个函数调用，则这个函数调用的参数会被立即求出。
	defer fmt.Println(haha())
	print("go func")
}

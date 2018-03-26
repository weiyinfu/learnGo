package main

import "fmt"

/**
go不支持函数重载
*/
func demo45haha() {
	fmt.Println("haha")
}

//func demo45haha(_ int) {
//	fmt.Println("haha")
//}
func main() {
	demo45haha()
	//demo45haha(3)
}

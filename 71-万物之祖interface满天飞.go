package main

import "fmt"

func print(x interface{}) {
	fmt.Println(x)
}
func main() {
	print(3)
	print("haha")
}

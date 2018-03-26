package main

import "fmt"

func haha(prefix string) {
	for i := 0; i < 3; i++ {
		fmt.Println(prefix, i)
	}
}
func main() {
	haha("one")
	go haha("baga")
	go func(prefix string) {
		for i := 0; i < 3; i++ {
			fmt.Println(prefix, i)
		}
	}("why")

	//如果这里不停一下,各个routine就会停止
	fmt.Scanln()
	fmt.Println("over")
}

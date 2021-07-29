package main

import "fmt"

func main() {
	haha := func(prefix string) {
		for i := 0; i < 3; i++ {
			fmt.Println(prefix, i)
		}
	}
	haha("one")
	go haha("baga")
	go func(prefix string) {
		for i := 0; i < 3; i++ {
			fmt.Println(prefix, i)
		}
	}("why")

	//如果这里不停一下,各个routine就会停止，这说明协程默认是deamon=True，是守护线程
	_, _ = fmt.Scanln()
	fmt.Println("over")
}

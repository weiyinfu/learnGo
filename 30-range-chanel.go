package main

import "fmt"

func main() {
	//如果此处q的长度等于1，那么在执行q<"two"的时候就会等待向q中放入数据，但是q中只有消费之后才能放入，因此形成死锁
	q := make(chan string, 2)
	q <- "one"
	q <- "two"
	close(q)
	for i := range q {
		fmt.Println(i)
	}
}

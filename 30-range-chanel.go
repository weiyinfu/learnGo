package main

import "fmt"

func main() {
	q := make(chan string, 2)
	q <- "one"
	q <- "two"
	close(q)
	for i := range q {
		fmt.Println(i)
	}
}

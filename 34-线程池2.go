package main

import "fmt"

func workerThread(id int, jobs <-chan int, result chan<- int) {
	for i := range jobs {
		result <- i * 2
	}
}
func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	for w := 0; w < 100; w++ {
		//这些worker完成的顺序不一定
		go workerThread(w, jobs, results)
	}
	for i := 0; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)
	for i := 0; i <= 5; i++ {
		fmt.Println(<-results)
	}
}

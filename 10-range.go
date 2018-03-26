package main

import "fmt"

func main() {
	//number是下标
	for number := range []int{4, 5, 6} {
		fmt.Println(number)
	}
	for index, number := range []int{1, 2, 3, 4} {
		fmt.Println(index, number)
	}
	for k, v := range (map[string]string{"username": "weidiao"}) {
		fmt.Println(k, "=", v)
	}
}

package main

import "fmt"

func main() {
	i := 1
	for i < 7 {
		fmt.Print(i)
		i++
	}
	for j := 7; j < 9; j++ {
		fmt.Print(i)
	}
	for {
		fmt.Println("loop forever")
		break
	}
	for i = 0; i < 7; i++ {
		if i < 4 {
			continue
		} else {
			fmt.Println("I am breaking")
			break
		}
	}
	fmt.Println("=======")
	var a = []int{1, 2, 3}
	for i := range a { //i是下标
		fmt.Println(i)
	}
	fmt.Println("=======")
	for i, v := range a {
		fmt.Println(i, v)
	}
	var ma = map[string]int{
		"one": 1,
		"two": 2,
	}
	for k, v := range ma {
		fmt.Println(k, v)
	}

}

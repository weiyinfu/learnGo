package main

import (
	"time"
	"fmt"
)

func main() {
	switch time.Now().Weekday() {
	case time.Monday:
		{
			fmt.Print("今天星期一")
		}
	case time.Tuesday:
		fmt.Println("今天星期二")
	default:
		fmt.Println(time.Now().Weekday())
	}
	whoami := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("I am bool")
		case int:
			fmt.Println("I am int")
		default:
			fmt.Println("who ma I?")
		}
	}
	whoami(true)
	whoami(3)
	//没有条件的switch
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

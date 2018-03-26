package main

import "fmt"

func main()  {
	i:=1
	for i<7{
		fmt.Print(i)
		i++
	}
	for j:=7;j<9;j++{
		fmt.Print(i)
	}
	for{
		fmt.Println("loop forever")
		break
	}
	for i=0;i<7;i++{
		if i<4{
			continue
		}else{
			fmt.Println("I am breaking")
			break
		}
	}
}

package main

import (
	"sort"
	"fmt"
)

func main()  {
	a:=[]int{1,4,6,5,2,3,7}
	sort.Ints(a)
	fmt.Println(a,sort.IntsAreSorted(a))
}

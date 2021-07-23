package main

import "fmt"

func main() {
	type Node struct {
		id int
	}
	nodeList := make([]Node, 0)
	a := map[int]*Node{}
	for i := 0; i < 10; i++ {
		nodeList = append(nodeList, Node{id: i})
	}
	for _, i := range nodeList {
		a[i.id] = &i
	}
	for k, v := range a {
		//最终的结果都是9
		fmt.Println(k, v.id)
	}
}

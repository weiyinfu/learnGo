package main

import (
	"fmt"
)

type Node2 struct {
	Value interface{}
	next  *Node2
}
type LinkedList2 struct {
	head Node2
}

func (list *LinkedList2) add(x *Node2) *Node2 {
	now := &list.head
	for ; now.next != nil; now = now.next {

	}
	now.next = x
	return x
}
func (self *LinkedList2) print() {
	for i := self.head.next; i != nil; i = i.next {
		fmt.Print(i.Value, " ")
	}
}
func main() {
	list := LinkedList2{}
	for i := 0; i < 10; i++ {
		list.add(&Node2{Value: i})
	}
	list.print()
}

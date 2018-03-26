package main

import "fmt"

type Node struct {
	Value interface{}
	next  *Node
}
type LinkedList struct {
	head Node
	tail Node
}

func (list *LinkedList) add(x *Node) *Node {
	now := &list.head
	for ; now.next != nil; now = now.next {

	}
	now.next = x
	return x
}
func (self *LinkedList) print() {
	for i := self.head.next; i != nil; i = i.next {
		fmt.Print(i.Value, " ")
	}
}
func main() {
	list := LinkedList{}
	for i := 0; i < 10; i++ {
		list.add(&Node{Value: i})
	}
	list.print()
}

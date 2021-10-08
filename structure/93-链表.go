package main

import (
	"container/list"
	"fmt"
	"reflect"
)

func main() {
	type Node struct {
		name string
	}
	a := list.New()
	a.PushBack(Node{"one"})
	a.PushBack(Node{"two"})
	a.PushFront("three")
	for i := a.Front(); i != nil; i = i.Next() {
		v := i.Value
		if reflect.TypeOf(v) == reflect.TypeOf(Node{}) {
			n := v.(Node)
			fmt.Println("node", n.name)
		} else {
			fmt.Println(v)
		}
	}
}

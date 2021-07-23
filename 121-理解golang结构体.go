package main

import (
	"fmt"
	"unsafe"
)

type A struct {
	name string
	x    int32
}
type B struct {
	name string
	x    int32
	y    int32
}
/**
A结构体的name字段是string，string类型是地址+长度，8B+8B，x是int32，4B，为了内存对齐，整个结构体占24B
B结构体string占16B，x和y占8B，整个结构体占24B
*/
func main() {
	a := A{name: "xx"}
	fmt.Println(unsafe.Sizeof(a))
	a = A{name: "xxxxxxxxxxxxx"}
	fmt.Println(unsafe.Sizeof(a))
	b := B{name: "xx"}
	fmt.Println(unsafe.Sizeof(b))
}

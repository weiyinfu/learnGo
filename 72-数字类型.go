package main

import "fmt"

/**

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8 的别名

rune // int32 的别名
     // 表示一个 Unicode 码点

float32 float64

complex64 complex128

*/
func main() {
	var nint8 int8 = 0
	var nint16 int16 = 0
	var ncomplex complex64 = complex(1, 2)
	ncomplex2 := 1 + 2i
	var nfloat32 float32 = 3.12
	var p = uintptr(nint16)
	fmt.Println(nint8, nint16, ncomplex, nfloat32, p, ncomplex2)
}

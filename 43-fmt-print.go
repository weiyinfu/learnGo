package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)

	fmt.Printf("%T\n", p)
	fmt.Printf("%t\n", true)

	//十进制
	fmt.Printf("%d\n", 123)
	//二进制
	fmt.Printf("%b\n", 123)
	//十六进制
	fmt.Printf("%x\n", 1234)
	//字符
	fmt.Printf("%c\n", 123)

	//浮点
	fmt.Printf("%f\n", 1.234)
	//指数形式
	fmt.Printf("%e\n", 12342423434.0)
	fmt.Printf("%E\n", 12342423434.0)

	fmt.Printf("\"%s\"\n", "weidiao")
	fmt.Printf("%q\n", "haha")
	fmt.Printf("%x\n", "haha")

	fmt.Printf("%p\n", &p)

	fmt.Printf("|%6d|%6d|\n", 234, 12)
	fmt.Printf("|%6.2f|%-6.2f|\n", 1.2, 3.4)
}

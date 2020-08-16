package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println(s)

	s[0] = "wei"
	//append的时候，发现数组长度不够，会新创建一个对象
	ss := append(s, "diao")
	fmt.Println(ss)
	ss[0] = "ha"
	fmt.Println(s)

	copy(s, ss)
	fmt.Println(s)
	fmt.Println(s[:1])
	//这说明go语言中的切片跟python中的切片不一样,go中的切片没有复制数组
	//但是go中的切片会自动扩张
	s[:1][0] = "baga"
	fmt.Println(s) //print baga

	//变长二维数组
	jiujiu := make([][]string, 9)
	for i := 1; i <= 9; i++ {
		jiujiu[i-1] = make([]string, i)
		for j := 1; j <= i; j++ {
			jiujiu[i-1][j-1] = fmt.Sprintf("%d*%d=%d", j, i, i*j)
		}
	}
	fmt.Println(jiujiu)
}

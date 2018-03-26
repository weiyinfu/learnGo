package main

import "fmt"

var username string
/**
每个文件被第一次访问的时候会调用init函数
*/
func init() {
	username = "weidiao"
}
func main() {
	fmt.Println(username)
}

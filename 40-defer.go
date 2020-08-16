package main

import (
	"fmt"
	"os"
)

func createFile(filename string) *os.File {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	return f
}
func closeFile(file *os.File) {
	fmt.Println("close")
	file.Close()
}
func write(file *os.File) {
	fmt.Println("write")
	fmt.Fprintln(file, "haha")
}
func knowDefer() {
	fmt.Println()
	conn := "connection"
	statement := "statement"
	fmt.Println("get", conn)
	fmt.Println("get", statement)
	defer fmt.Println("close", conn)      //这句话会后执行
	defer fmt.Println("close", statement) //这句话会先执行
	fmt.Println("读取数据库")
	fmt.Println()
}
func deferStack() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
func main() {
	f := createFile("haha.txt")
	defer closeFile(f)
	write(f)
	defer fmt.Println("这句话会在主函数执行完成之后打印")
	fmt.Println("这句话会先打印")
	knowDefer()
	deferStack()
}

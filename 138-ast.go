package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

/*
一个golang结构可视化网站
http://goast.yuroyoro.net/
*/
func main() {
	const source = `
 package main
 func main() {
 answer := 42
for i:=0;i<10;i++{
answer=i
}
 }`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "", source, parser.AllErrors)
	ast.Print(fset, file)

	//解析结果
	expr := &ast.BinaryExpr{
		X: &ast.BasicLit{
			Value: "3",
			Kind:  token.INT,
		},
		Op: token.ADD,
		Y: &ast.BasicLit{
			Value: "2",
			Kind:  token.INT,
		},
	}
	fmt.Println(expr)
}

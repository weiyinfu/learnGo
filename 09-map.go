package main

import "fmt"

func main()  {
	m:=make(map[string]string)
	m["username"]="weidiao"
	m["age"]="haha"
	fmt.Println(m)
	delete(m,"username")
	fmt.Println(m)
	_,what:=m["age"]
	fmt.Println(what)
	fmt.Println(map[string]string{
		"name":"weidiao",
	})
}

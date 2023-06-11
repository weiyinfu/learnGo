package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id   int64  `json:"id,string"` //使用这个注释可以将int64转成string
	Name string `json:"name"`
}

func main() {
	x := Student{Id: 2134, Name: "weiyinfu"}
	s, _ := json.Marshal(x)
	fmt.Println(string(s))
	s = []byte(`{"id":2134,"name":"weiyinfu"}`)
	e := json.Unmarshal(s, &x)
	if e != nil {
		fmt.Println(e)
		return
	}
	s, _ = json.Marshal(x)
	fmt.Println(string(s))
}

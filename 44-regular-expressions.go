package main

import (
	"regexp"
	"fmt"
	"bytes"
)

func main() {
	res, err := regexp.MatchString("p[a-z]*ch", "peach")
	fmt.Println(res, err)

	//compile & match
	pattern, err := regexp.Compile("p[a-z]*ch")
	fmt.Println(pattern.MatchString("peach"), err)
	fmt.Println(pattern.FindString("I love eating peach"))
	fmt.Println(pattern.FindStringIndex("I love eating peach"))
	fmt.Println(pattern.FindStringSubmatch("I love eating peach"))
	fmt.Println(pattern.FindStringSubmatchIndex("I love eating peach for peach is delicious"))
	fmt.Println(pattern.FindAllString("I love eating peach for peach is delicious", -1))
	fmt.Println(pattern.Match([]byte("peach")))
	fmt.Println(pattern.ReplaceAllString("I love eating peach for peach is delicious", "<fruit>"))

	in := []byte("i like peach haha")
	out := pattern.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}

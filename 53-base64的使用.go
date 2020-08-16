package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := base64.StdEncoding.EncodeToString([]byte("weidiao"))
	fmt.Println(s)
	res, err := base64.StdEncoding.DecodeString(s)
	fmt.Println(string(res), err)

	fmt.Println(base64.URLEncoding.EncodeToString([]byte("weidiao")))
}

package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	a := sha1.New()
	a.Write([]byte("weidiao"))
	fmt.Printf("%x\n", a.Sum(nil))
	fmt.Println(sha1.Sum([]byte("weidiao")))
}

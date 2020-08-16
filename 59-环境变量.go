package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))
	//Use os.Environ to list all key/value pairs in the environment. This returns a slice of strings in the form KEY=value. You can strings.Split them to get the key and value. Here we print all the keys.

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0], "=", pair[1])
	}
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	var p = fmt.Println
	p(strings.Compare("weidiao", "haha"))
	p(strings.Contains("weidiao", "idi"))
	p(strings.Count("weidiao", "i"))
	p(strings.HasPrefix("weidiao", "wei"))
}

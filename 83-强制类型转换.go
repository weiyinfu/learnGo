package main

import (
	"fmt"
	"strconv"
)

func haha2(s interface{}, ty string) interface{} {
	switch ty {
	case "int":
		value, _ := strconv.Atoi(s.(string))
		return value
	case "int64":
		value, _ := strconv.ParseInt(s.(string), 10, 64)
		return value
	case "string":
		return s
	default:
		return nil
	}
}
func main() {
	var x int64
	x = haha2("3", "int64").(int64)
	fmt.Println(x)
}

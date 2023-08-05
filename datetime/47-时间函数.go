package main

import (
	"fmt"
	"time"
)

func main() {
	print := fmt.Println
	print(time.Now())
	then := time.Date(2018, 3, 20, 22, 20, 3, 234, time.Local)
	print(then, then.Year(), then.Month(), then.Second(), then.Weekday())

	print(then.Before(time.Now()), then.After(time.Now()), time.Now().Sub(then))

	print(time.Now().Sub(then).Hours())
}

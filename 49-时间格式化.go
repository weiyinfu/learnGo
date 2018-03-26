package main

import (
	"fmt"
	"time"
)

func main() {
	var print = fmt.Println
	t := time.Now()
	print(t.Format(time.RFC3339))

	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	print(t1, e)

	/**
	go语言中的format真是设计的高级,直接让你写一个例子
	*/
	print(t.Format("3:04PM"))
	print(t.Format("Mon Jan _2 15:04:05 2006"))
	print(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	print(t2)


	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	res, e := time.Parse(ansic, "8:41PM")
	print(res,"haha\nbaga",e)
}

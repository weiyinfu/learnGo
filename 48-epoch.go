package main

import (
	"time"
	"fmt"
)

func main() {
	print := fmt.Println
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	print(now.Second(), now.Unix())
	print(now.Nanosecond(), now.UnixNano())
	print(secs, nanos)
	print(time.Unix(now.Unix(),now.UnixNano()))
}

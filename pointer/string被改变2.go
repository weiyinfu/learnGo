package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type Monitor struct {
	Region string
}

func solve(c context.Context, region string) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println(err)
				}
			}()
			time.Sleep(time.Second * 3)
			fmt.Println(fmt.Sprintf("%v %v", region, i))
		}(i)
	}
}
func main() {
	ctx := context.Background()
	region := "cn"
	solve(ctx, region)
	s, _ := json.Marshal("baga")
	_ = json.Unmarshal(s, &region)
	time.Sleep(time.Second * 10)
}

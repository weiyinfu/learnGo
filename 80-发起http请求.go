package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main() {
	resp, _ := http.Get("http://cnblogs.com/weiyinfu")
	defer resp.Body.Close()
	res, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(res))
}

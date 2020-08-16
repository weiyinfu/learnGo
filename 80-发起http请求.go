package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
go get github.com/PuerkitoBio/goquery
*/
func main() {
	resp, _ := http.Get("http://cnblogs.com/weiyinfu")
	defer resp.Body.Close()
	res, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(res))
}

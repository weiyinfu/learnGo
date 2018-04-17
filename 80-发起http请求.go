package main

import (
	"net/http"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"bytes"
	"strings"
)

/**
go get github.com/PuerkitoBio/goquery
*/
func main() {
	resp, _ := http.Get("http://cnblogs.com/weiyinfu")
	defer resp.Body.Close()
	res, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(res))
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(res))
	doc.Find(".postTitle").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(strings.Trim(selection.Text(), "\n"))
	})
}

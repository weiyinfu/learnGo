package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

var datas []string

func Add(str string) string {
	data := []byte(str)
	sData := string(data)
	datas = append(datas, sData)

	return sData
}
func main() {
	go func() {
		for {
			time.Sleep(time.Millisecond * 100)
			log.Println(Add("https://github.com/EDDYCJY"))
		}
	}()
	fmt.Println("please visit http://127.0.0.1:6060/debug/pprof/")
	http.ListenAndServe("0.0.0.0:6060", nil)
}

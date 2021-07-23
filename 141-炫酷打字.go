package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	road := make(chan rune)
	done := make(chan bool)
	allDone := make(chan bool)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println("读协程失败")
				panic(err)
			}
		}()
		filename := "why.txt"
		f, e := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
		if e != nil {
			panic(fmt.Errorf("打开文件%v失败err=%v", filename, e))
		}
		cin := bufio.NewReader(f)
		for {
			ch, _, err := cin.ReadRune()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println(err)
				panic("读取文件失败")
			}
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			road <- ch
		}
		done <- true
	}()
	go func() {
		for {
			select {
			case ch := <-road:
				{
					fmt.Print(string(ch))
				}
			case <-done:
				{
					fmt.Println("完事了")
					allDone <- true
				}
			}
		}
	}()
	<-allDone
}

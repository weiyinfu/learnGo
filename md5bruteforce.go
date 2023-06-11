package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func md5v(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
func main() {
	if len(os.Args) == 1 {
		fmt.Println(`
usage: command captcha
`)
		os.Exit(0)
	}
	candidates := []string{}
	for i := 0; i < 10; i++ {
		candidates = append(candidates, strconv.Itoa(i))
	}
	for i := 0; i < 26; i++ {
		candidates = append(candidates, string(rune('a'+i)))
		candidates = append(candidates, string(rune('A'+i)))
	}
	solve(candidates)
	//fmt.Println(md5v("00vKcX"))
}
func solve(candidates []string) {
	prefix := strings.TrimSpace(os.Args[1])
	for _, i := range candidates {
		for _, j := range candidates {
			for _, k := range candidates {
				for _, l := range candidates {
					for _, m := range candidates {
						for _, n := range candidates {
							plain := i + j + k + l + m + n
							s := md5v(plain)
							if strings.HasPrefix(s, prefix) {
								fmt.Println(plain, s, prefix)
								goto out
							}
						}
					}
				}
			}
		}
	}
out:
}

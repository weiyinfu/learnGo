package main

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"math"
)

func EncryptAES(key string, plain string) string {
	//key的长度应该为16B，超出无效
	keyBytes := make([]byte, 16)
	copy(keyBytes, key)
	plainBytes := make([]byte, int(math.Ceil(float64(len(plain))/16.0)*16))
	copy(plainBytes, plain)
	c, _ := aes.NewCipher(keyBytes)
	ans := make([]byte, len(plainBytes))
	for i := 0; i < len(plainBytes); i += 16 {
		c.Encrypt(ans[i:], plainBytes[i:])
	}
	return base64.StdEncoding.EncodeToString(ans)
}
func main() {
	plaintext := "xadfasdf;xxxx"
	//gfWwACXHxBO5qwTEqs/fWgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
	key := "a6f8hn3jhud7sdfjh2hh44672h"
	encoded := EncryptAES(key, plaintext)
	fmt.Println(encoded)
}

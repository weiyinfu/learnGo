package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"strings"
)

// Sign 计算签名值, ver有两种 auth-v1、auth-v2，参见文档：https://bytedance.feishu.cn/wiki/wikcneifTKfIHZ60UUDnq9xvoJg#
func Sign(ver, ak, sk, now, expire string, body []byte) string {
	signKeyInfo := fmt.Sprintf("%s/%s/%s/%s", ver, ak, now, expire)
	signKey := sha256HMAC([]byte(sk), []byte(signKeyInfo))
	return fmt.Sprintf("%s/%s", signKeyInfo, string(sha256HMAC(signKey, body)))
}
func sha256HMAC(key, data []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(data)
	return []byte(fmt.Sprintf("%x", mac.Sum(nil)))
}
func main() {
	body := "{\"size\":50,\"next_id\":0,\"version\":0,\"open_id\":\"841bd637-1229-44ef-ac0a-6d3e43b5cf3e\",\"force_pull\":true}"
	xx := "asdfasdf"
	ks := strings.Split(xx, "/")

	ans := Sign(ks[0], ks[1], "asdfasdfas", ks[2], ks[3], []byte(body))
	fmt.Println(ans)
	fmt.Println(ans == xx)
}

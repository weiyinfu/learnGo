package main

import "fmt"

func index(a []string, s string) int {
	ind := -1
	for ind, x := range a {
		if s == x {
			return ind
		}
	}
	return ind
}
func include(a []string, s string) bool {
	return index(a, s) != -1
}
func any(a []string, s string) bool {
	return index(a, s) != -1
}
func all(a []string, s string) bool {
	for _, x := range a {
		if x != s {
			return false
		}
	}
	return true
}
func filter(a []string, f func(string) bool) []string {
	ans := []string{}
	for _, x := range a {
		if f(x) {
			ans = append(ans, x)
		}
	}
	return ans
}
func map_(a []string, f func(string) string) []string {
	ans := make([]string, len(a))
	for ind, x := range a {
		ans[ind] = f(x)
	}
	return ans
}
func main() {
	a := []string{"we", "idi", "ao"}
	fmt.Println(index(a, "idi"))
	fmt.Println(all(a, "idi"))
	fmt.Println(any(a, "idi"))
	fmt.Println(filter(a, func(s string) bool {
		return s > "haha"
	}))
	fmt.Println(map_(a, func(s string) string {
		return s + "haha"
	}))
}

package main

import (
	"math/rand"
	"sort"
)

type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	a := make([]int, 10)
	println(len(a))
	for i := 0; i < 10; i++ {
		a = append(a, rand.Int())
	}
	sort.Sort(IntSlice(a))
}

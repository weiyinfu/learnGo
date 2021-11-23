package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func generateProblem() []int {
	//生成一个不包含重复元素的数组
	var a []int
	for i := 0; i < 100; i++ {
		a = append(a, i)
	}
	rand.Shuffle(100, func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}
func bruteforce(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[i] {
				s++
			}
		}
	}
	return s
}

func mergeSort2_(a []int, b []int, beg, end int) int {
	if beg >= end {
		return 0
	}
	mid := (beg + end) / 2
	s := 0
	s += mergeSort2_(a, b, beg, mid)
	s += mergeSort2_(a, b, mid+1, end)
	i, j := beg, mid+1
	k := beg
	for k <= end {
		var use1 bool
		if i <= mid && j <= end {
			if a[i] < a[j] {
				use1 = true
			} else {
				use1 = false
			}
		} else if i > mid {
			use1 = false
		} else if j > end {
			use1 = true
		}
		if use1 {
			b[k] = a[i]
			k++
			i++
		} else {
			s += j - k
			b[k] = a[j]
			k++
			j++
		}
	}
	for i := beg; i <= end; i++ {
		a[i] = b[i]
	}
	return s
}
func mergeSort2(a []int) int {
	b := make([]int, len(a))
	return mergeSort2_(a, b, 0, len(a)-1)
}
func lowbit(x int) int {
	return x & (-x)
}
func treeArray_update(tr []int, x int) {
	i := x
	for i < len(tr) {
		tr[i]++
		i += lowbit(i)
	}
}
func treeArray_query(tr []int, x int) int {
	i := x
	s := 0
	for i > 0 {
		s += tr[i]
		i -= lowbit(i)
	}
	return s
}
func treeArray(a []int) int {
	//将数字离散化
	numMap := map[int]int{}
	b := make([]int, len(a))
	copy(b, a)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	for ind, v := range b {
		numMap[v] = ind
	}
	for ind, v := range a {
		b[ind] = numMap[v]
	}
	//离散化结束
	tr := make([]int, 1+len(a))
	s := 0
	for ind, v := range b {
		small := treeArray_query(tr, v+1) //求小于v的数字的个数
		s += ind - small                  //比v+1大的数字的个数
		treeArray_update(tr, v+1)         //让大于v+1的数字的个数加上1
	}
	return s
}
func main() {
	a := generateProblem()
	for _, x := range []struct {
		name string
		f    func([]int) int
	}{
		{"bruteforce", bruteforce},
		{"mergesort", mergeSort2},
		{"treeArray", treeArray},
	} {
		b := make([]int, len(a))
		copy(b, a)
		ans := x.f(b)
		fmt.Println(x.name, ans)
	}
}

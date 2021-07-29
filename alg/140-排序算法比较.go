package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

func deepCopy(a []int) []int {
	b := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		b[i] = a[i]
	}
	return b
}
func GenerateProblem(n int, m int) ([]int, []int) {
	//产生一个问题，返回问题和答案
	a := make([]int, n)
	for i := 0; i < n; i++ {
		if m == 0 {
			a[i] = rand.Int()
		} else {
			a[i] = rand.Int() % m
		}
	}
	b := deepCopy(a)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return a, b
}
func selectSort(a []int) {
	for i := 0; i < len(a); i++ {
		small := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[small] {
				small = j
			}
		}
		a[i], a[small] = a[small], a[i]
	}
}
func swapSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[i] {
				a[j], a[i] = a[i], a[j]
			}
		}
	}
}
func bubbleSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := len(a) - 1; j > i; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}
func insertSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			} else {
				break
			}
		}
	}
}
func quickSort_(a []int, beg, end int) {
	if beg >= end {
		return
	}
	l, r := beg, end
	mid := a[l]
	for l < r {
		for r > l && a[r] >= mid {
			r--
		}
		if l == r {
			break
		}
		a[l] = a[r]
		a[r] = mid
		l++
		for r > l && a[l] <= mid {
			l++
		}
		if l == r {
			break
		}
		a[r] = a[l]
		a[l] = mid
		r--
	}
	quickSort_(a, beg, l-1)
	quickSort_(a, l+1, end)
}
func quickSort(a []int) {
	quickSort_(a, 0, len(a)-1)
}
func mergeSort_(a []int, b []int, beg, end int) {
	if beg >= end {
		return
	}
	mid := (beg + end) / 2
	mergeSort_(a, b, beg, mid)
	mergeSort_(a, b, mid+1, end)
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
			b[k] = a[j]
			k++
			j++
		}
	}
	for i := beg; i <= end; i++ {
		a[i] = b[i]
	}
}
func mergeSort(a []int) {
	b := make([]int, len(a))
	mergeSort_(a, b, 0, len(a)-1)
}
func same(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

type sortMethod func([]int)

func main() {
	ma := map[string]sortMethod{
		"冒泡": bubbleSort,
		"选择": selectSort,
		"插入": insertSort,
		"快排": quickSort,
		"交换": swapSort,
		"归并": mergeSort,
	}
	testTimes := 10000
	n := 800
	timeUsedMap := sync.Map{}
	wg := sync.WaitGroup{}
	for name, f := range ma {
		wg.Add(1)
		go func(name string, f sortMethod) {
			defer func() {
				wg.Done()
			}()
			beginTime := time.Now()
			for testI := 0; testI < testTimes; testI++ {
				x, y := GenerateProblem(n, 0)
				yy := deepCopy(x)
				f(yy)
				if !same(y, yy) {
					fmt.Printf("%v方法有瑕疵", name)
					panic("done")
				}
			}
			endTime := time.Now()
			timeUsed := endTime.Sub(beginTime)
			timeUsedMap.Store(name, timeUsed)
		}(name, f)
	}
	wg.Wait()
	timeUsedMap.Range(func(key, value interface{}) bool {
		fmt.Println(key, value.(time.Duration).String())
		return true
	})
}

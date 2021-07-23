package main

import (
	"fmt"
	"sort"
)

type Student struct {
	score int
}

type StudentList []Student

func (a StudentList) Len() int {
	return len(a)
}
func (a StudentList) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a StudentList) Less(i, j int) bool {
	return a[i].score < a[j].score
}
func main() {
	a := []int{1, 4, 6, 5, 2, 3, 7}
	sort.Ints(a)
	fmt.Println(a, sort.IntsAreSorted(a))
	students := StudentList{{1}, {2}, {3}}
	b := make([]Student, 10)
	copy(b, students) //copy只是简单的copy
	sort.Sort(students)
	fmt.Println(students)
	sort.Slice(b, func(i, j int) bool {
		return b[i].score < b[j].score
	})
	fmt.Println(b)
}

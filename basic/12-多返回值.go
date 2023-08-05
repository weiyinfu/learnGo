package main

import (
	"fmt"
	"math"
)

func maxMin(a []int) (float64, float64) {
	var mi float64
	var ma float64
	for _, x := range a {
		mi = math.Min(mi, float64(x))
		ma = math.Max(ma, float64(x))
	}
	return ma, mi
}
func main() {
	fmt.Print(maxMin([]int{1, 2, 3, 4, 5}))
}

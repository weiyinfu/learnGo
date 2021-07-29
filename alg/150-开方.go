package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	"unsafe"
)

func BiSearch(x, delta float64) float64 {
	l := 0.0
	r := x
	if x < 1 {
		r = 1
	}
	for l+delta < r {
		mid := (l + r) / 2
		if mid*mid > x {
			r = mid
		} else {
			l = mid
		}
	}
	return l
}
func magicNumber(x, delta float64) float64 {
	solve := func(x float32) float32 {
		threehalfs := float32(1.5)
		x2 := x * 0.5
		y := x
		i := (*int32)(unsafe.Pointer(&y)) // evil floating point bit level hacking
		*i = 0x5f3759df - (*i >> 1)       // what the fuck?
		yy := (*float32)(unsafe.Pointer(i))
		z := *yy
		z = z * (threehalfs - (x2 * z * z)) // 1st iteration
		return 1 / z
	}
	return float64(solve(float32(x)))
}
func newton(x float64, delta float64) float64 {
	z := x
	y2 := x
	for {
		y1 := z * z
		t := (y2-y1)/(2*z) + z
		if math.Abs(z-t) < delta {
			break
		}
		z = t
	}
	return z
}
func StadardLib(x float64, delta float64) float64 {
	/**
	此代码copy自golang官方库，golang正版实现是汇编，并非此代码
	*/
	const (
		mask  = 0x7FF
		shift = 64 - 11 - 1
		bias  = 1023
	)

	// special cases
	switch {
	case x == 0 || math.IsNaN(x) || math.IsInf(x, 1):
		return x
	case x < 0:
		return math.NaN()
	}
	ix := math.Float64bits(x)
	// normalize x
	exp := int((ix >> shift) & mask)
	if exp == 0 { // subnormal x
		for ix&(1<<shift) == 0 {
			ix <<= 1
			exp--
		}
		exp++
	}
	exp -= bias // unbias exponent
	ix &^= mask << shift
	ix |= 1 << shift
	if exp&1 == 1 { // odd exp, double x to make it even
		ix <<= 1
	}
	exp >>= 1 // exp = exp/2, exponent of square root
	// generate sqrt(x) bit by bit
	ix <<= 1
	var q, s uint64               // q = sqrt(x)
	r := uint64(1 << (shift + 1)) // r = moving bit from MSB to LSB
	for r != 0 {
		t := s + r
		if t <= ix {
			s = t + r
			ix -= t
			q += r
		}
		ix <<= 1
		r >>= 1
	}
	// final rounding
	if ix != 0 { // remainder, result not exact
		q += q & 1 // round according to extra bit
	}
	ix = q>>1 + uint64(exp-1+bias)<<shift // significand + biased exponent
	return math.Float64frombits(ix)
}
func testOne() {
	x := 3.0
	delta := 1e-6
	y := BiSearch(x, delta)
	fmt.Println("答案", math.Sqrt(x))
	fmt.Println("二分法", y)
	fmt.Println("魔数法", magicNumber(x, delta))
	fmt.Println("牛顿法", newton(x, delta))
	fmt.Println("官方库2", StadardLib(x, delta))
}
func testMany() {
	ta := map[string]func(x, delta float64) float64{
		"官方库": func(x, delta float64) float64 {
			return math.Sqrt(x)
		},
		"二分法":     BiSearch,
		"魔数法":     magicNumber,
		"牛顿法":     newton,
		"官方库具体实现": StadardLib,
	}
	times := 100000
	for name, f := range ta {
		begin := time.Now()
		for i := 0; i < times; i++ {
			x := rand.Float64() * 100
			f(x, 1e-4)
		}
		end := time.Now()
		fmt.Println(name, end.Sub(begin).String())
	}
}
func main() {
	testOne()
	fmt.Println("==========")
	testMany()
}

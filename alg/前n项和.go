package main

/**
求1^k+2^k+...+n^k

当n远远大于k的时候，使用数学方法比较快，否则不如暴力简单。

公式法就像活字印刷术，如果只算一次，并不快，如果计算很多次，就会变得非常快。
*/
import (
	"fmt"
	"time"
)

const BIG = int64(1e9 + 7)

func pow(x, k int) int64 {
	if k == 0 {
		return 1
	}
	y := pow(x, k/2)
	ans := y * y
	ans %= BIG
	if k%2 == 1 {
		ans *= int64(x)
		ans %= BIG
	}
	return ans
}
func brute(n, k int) int {
	s := int64(0)
	for i := 1; i <= n; i++ {
		s += pow(i, k)
		s %= BIG
	}
	return int(s)
}
func rev(x int) int64 {
	//求逆
	var a, b int64
	_ = gcd(BIG, int64(x), &a, &b)
	if b < 0 {
		b += BIG
	}
	return b
}

func gcd(x int64, y int64, a *int64, b *int64) int64 {
	if y == 0 {
		*a = 1
		*b = 1
		return x
	}
	d := gcd(y, x%y, b, a)
	*b -= x / y * *a
	*a %= BIG
	*b %= BIG
	return d
}
func C(n, k int) int64 {
	ans := int64(1)
	for i := 1; i <= k; i++ {
		ans *= int64(n + 1 - i)
		ans %= BIG
		ans *= rev(i)
		ans %= BIG
	}
	return ans
}
func smart(n, k int) int {
	S := make([]int64, k+1)
	S[0] = int64(n)
	for i := 1; i <= k; i++ {
		up := pow(n+1, i+1) - 1
		for j := 0; j < i; j++ {
			up -= C(i+1, j) * S[j]
			up %= BIG
		}
		S[i] = up * rev(i+1)
		S[i] %= BIG
		if S[i] < 0 {
			S[i] += BIG
		}
	}
	return int(S[k])
}
func buildFormula(k int) [][]int64 {
	S := make([][]int64, k+1)
	for i := 0; i < len(S); i++ {
		S[i] = make([]int64, i+1)
	}
	S[0][0] = 1
	//杨辉三角预先处理组合数
	C := make([][]int64, k+2)
	for i := 0; i < len(C); i++ {
		C[i] = make([]int64, k+2)
		C[i][0] = 1
	}
	for i := 1; i < len(C); i++ {
		for j := 1; j <= i; j++ {
			C[i][j] = C[i-1][j] + C[i-1][j-1]
			C[i][j] %= BIG
		}
	}
	for i := 1; i <= k; i++ {
		for k := 0; k < i; k++ {
			for j := 0; j <= k; j++ {
				S[i][j] -= C[i+1][k] * S[k][j]
				S[i][j] %= BIG
			}
		}
		R := rev(i + 1)
		for j := 0; j <= i; j++ {
			S[i][j] += C[i+1][j+1]
			S[i][j] *= R
			S[i][j] %= BIG
			if S[i][j] < 0 {
				S[i][j] += BIG
			}
		}
	}
	return S
}

// 使用公式法求解
func formula(n, k int) int {
	S := buildFormula(k)
	return int(apply(S[k], n))
}

func apply(form []int64, n int) int64 {
	s := int64(0)
	x := int64(1)
	for i := 0; i < len(form); i++ {
		x *= int64(n)
		x %= BIG
		s += x * form[i]
		s %= BIG
	}
	return s
}
func testGcd() {
	var a, b int64
	fmt.Println(gcd(12, 8, &a, &b))
}
func testRev() {
	r := rev(3)
	fmt.Println(r, 3*r%BIG)
	rr := rev(int(r))
	fmt.Println(rr)
}
func testRight() {
	n, k := 200000, 300
	begTime := time.Now().UnixMilli()
	realAns := brute(n, k)
	endTime := time.Now().UnixMilli()
	myAns := smart(n, k)
	endTime2 := time.Now().UnixMilli()
	formAns := formula(n, k)
	endTime3 := time.Now().UnixMilli()
	fmt.Println(realAns, myAns, formAns)
	fmt.Println("暴力耗时", endTime-begTime, "递推方法耗时", endTime2-endTime, "公式法耗时", endTime3-endTime2)
}
func main() {
	//for i := 2; i < 100; i++ {
	//	beg := time.Now().UnixMilli()
	//	buildFormula(i)
	//	end := time.Now().UnixMilli()
	//	fmt.Println(i, "time used", end-beg)
	//}
	testRight()
}

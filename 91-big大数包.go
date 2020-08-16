package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := big.NewInt(1)
	y := big.NewInt(2)
	z := big.NewInt(0)
	z.Mul(x, y).Add(x, y) //z=x*y;z=x+y
	fmt.Println(z)
}

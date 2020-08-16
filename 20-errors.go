package main

import (
	"fmt"
	"math/rand"
)

/**
GO语言中的error是一个接口
type error interface {
    Error() string
}
*/

type DivideZeroError struct {
	arg  int
	desc string
}

func (e DivideZeroError) Error() string {
	return fmt.Sprint(e.desc)
}
func div(x, y int) (int, error) {
	if y == 0 {
		return 0, DivideZeroError{x, "divide by zero"}
	} else {
		return x / y, nil
	}
}
func main() {
	for {
		x, y := rand.Int()%4-2, rand.Int()%4-2
		z, err := div(x, y)
		fmt.Println(x, y, z, err)
		fmt.Scanln()
	}

}

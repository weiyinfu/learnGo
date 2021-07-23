package main

import (
	"sync/atomic"
)

type One struct {
}

func (*One) Error() string {
	return "one"
}

type Two struct {
}

func (*Two) Error() string {
	return "two"
}
func main() {
	//atomic放东西必须需要放同样的东西
	x := atomic.Value{}
	var one, two error
	one = &One{}
	two = &Two{}
	x.Store(one)
	x.Store(two)
}

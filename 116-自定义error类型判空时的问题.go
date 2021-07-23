package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	name string
}

func (m *MyError) Error() string {
	return m.name
}
func myfunc() (string, *MyError) {
	return "", nil
}
func myfunc2() (string, error) {
	return "", nil
}
func main() {
	e := errors.New("why")
	_, e = myfunc()
	fmt.Println(e == nil, e == (*MyError)(nil), e)
	_, e = myfunc2()
	fmt.Println(e == nil, e == (*MyError)(nil), e)
}

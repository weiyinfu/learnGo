package main

import (
	"fmt"
	"reflect"
)

func main() {

	type Node struct {
		name string `weidiao` //这种注解被称为tag，是golang特有的语法，是结构体字段上的东西
		age  int    `baga`
	}
	x := Node{}
	x.name = "weiyinfu"
	x.age = 12
	xv := reflect.ValueOf(x)

	fmt.Println(xv)
	fmt.Println("Type:", xv.Type(),
		//"Addr", xv.Addr(),//报错，unaddressable
		"\nString", xv.String(),
		//"int", xv.Int(),//报错， call of reflect.Value.Int on struct Value
		//"bool", xv.Bool(),//panic: reflect: call of reflect.Value.Bool on struct Value
		//"bytes", xv.Bytes(), //panic: reflect: call of reflect.Value.Bytes on struct Value

		"\ncanAddr", xv.CanAddr(),
		"\ncanInterface", xv.CanInterface(),
		"\ncanSet", xv.CanSet(),
		"\nisValid", xv.IsValid(),
		"\nIsZero", xv.IsZero(),
		//"\nIsNil", xv.IsNil(), //panic: reflect: call of reflect.Value.IsNil on struct Value

		"\n==========\n",
		"\nnumField", xv.NumField(),
		"\nnumMethod", xv.NumMethod(),
		"\n",
	)
	xt := reflect.TypeOf(x)
	fmt.Println(
		"name", xt.Name(),
		"\nString", xt.String(),
	)
	for i := 0; i < xv.NumField(); i++ {
		f := xv.Field(i)
		name := xt.Field(i)
		fmt.Println("name=", name.Name,
			"type=", name.Type,
			"index=", name.Index,
			"tag=", name.Tag,
			"anonymous=", name.Anonymous,
			"offset=", name.Offset,
			"pkgPath=", name.PkgPath,
			"canInterface=", f.CanInterface())
		if f.Kind() == reflect.String {
			fmt.Println("value=", f.String())
		} else if f.Kind() == reflect.Int {
			fmt.Println("value=", f.Int())
		}
	}
}

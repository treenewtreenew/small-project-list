package main

import (
	"fmt"
	"io"
)

type MyInterface interface {
	M1(int) error
	M2(int, ...string)
}

type Interface1 interface {
	M1()
}
type Interface2 interface {
	// M1(string) // 编译器报错：duplicate method M1
	M11(string)
	M2()
}

type Interface3 interface {
	Interface1
	Interface2 // 编译器报错：duplicate method M1
	M3()
}

type EmptyInterface interface {
}

var err error
var r io.Reader

func foo1() {
	var i interface{} = 15
	i = "hello, golang"
	fmt.Println(i)
	type T struct{}
	var t T
	i = t
	i = &t

	v, ok := i.(T)
	fmt.Println(v, ok)

	v1 := i.(*T)
	fmt.Print(v1)

}

func foo2() {
	var a int64 = 13
	var i interface{} = a
	v1, ok := i.(int64)
	fmt.Printf("v1=%d, the type of v1 is %T, ok=%t\n", v1, v1, ok) // v1=13, the type of v1 is int64, ok=true
	v2, ok := i.(string)
	fmt.Printf("v2=%s, the type of v2 is %T, ok=%t\n", v2, v2, ok) // v2=, the type of v2 is string, ok=false
	v3 := i.(int64)
	fmt.Printf("v3=%d, the type of v3 is %T\n", v3, v3) // v3=13, the type of v3 is int64
	v4 := i.([]int)                                     // panic: interface conversion: interface {} is int64, not []int
	fmt.Printf("the type of v4 is %T\n", v4)
}

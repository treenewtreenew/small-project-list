package main

import (
	"fmt"
	"io"
	"net/http"
)

// receiver argument type cannot be pointer type or interface type
type MyInt *int

// invalid receiver type MyInt (pointer or interface type)c
func (r MyInt) String() string {
	return fmt.Sprintf("%d", *(*int)(r))
}

type MyReader io.Reader // the io.Reader is interface
// invalid receiver type MyReader (pointer or interface
func (r MyReader) Read(p[byte]) (int, error) {
	return r.Read(p)
}

// Go 要求，方法声明要与 receiver 参数的基类型声明放在同一个包内。
// 基于这个约束，我们还可以得到两个推论。
// 1st : 我们不能为原生类型（诸如 int、float64、map 等）添加方法。
// cannot define new methods on non-local type int
func (i int) Foo() string {
	return fmt.Sprintf("%d", i)
}

// cannot define new methods on non-local type http.Servercompiler
// 2nd : cannot declare new method for other type in other go package crossing packages
func (s http.Server) Foo() {

}

// the essence of method
type T struct {
	a int
}

func (t T) Get() int {
	return t.a
}

func (t *T) Set(a int) int {
	t.a = a
	return t.a
}

// equivalent
// 类型T的方法Get的等价函数
func Get(t T) int {
	return t.a
}

// 类型*T的方法Set的等价函数
func Set(t *T, a int) int {
	t.a = a
	return t.a
}

// method expression
func methodExpression() {
	var t T
	t.Get()
	(&t).Set(1)

	// equivalent
	T.Get(t)
	(*T).Set(&t, 1)
}

// essence : Go 语言中的方法的本质就是，一个以方法的 receiver 参数作为第一个参数的普通函数。

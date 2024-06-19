package main

import (
	"fmt"
	"io"
	"strings"
)

// type embedding
// 接口型类型嵌入
type E interface {
	M1()
	M2()
}

// example of interface type embedding
/*
type I interface {
	M1()
	M2()
	M3()
}
*/

type I interface {
	E
	M3()
}

// go source code example
// $GOROOT/src/io/io.go

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

type ReadWriter interface {
	Reader
	Writer
}

type ReadCloser interface {
	Reader
	Closer
}

type WriteCloser interface {
	Writer
	Closer
}

type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

// struct type embedding
type T1 int
type t2 struct {
	n int
	m int
}

type I1 interface {
	M1()
}

// struct type embedding T1 *t2 I1, Embedded Field
// 既代表字段的名字，也代表字段的类型
// 嵌入字段类型的底层类型不能为指针类型
type S1 struct {
	T1
	*t2
	I1
	a int
	b string
}

type MyInt int

func (n *MyInt) Add(m int) {
	*n = *n + MyInt(m)
}

type t struct {
	a int
	b int
}

type S struct {
	*MyInt
	t
	io.Reader
	s string
	n int
}

func main() {
	m := MyInt(17)
	r := strings.NewReader("hello, go")
	s := S{
		MyInt: &m,
		t: t{
			a: 1,
			b: 2,
		},
		Reader: r,
		s:      "demo",
	}

	var sl = make([]byte, len("hell"))
	s.Reader.Read(sl)
	sl = make([]byte, len("hello, go"))
	s.Read(sl)
	fmt.Println(string(sl))
	s.MyInt.Add(5)
	s.Add(3)
	fmt.Println(*(s.MyInt))
}

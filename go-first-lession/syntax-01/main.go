package main

import "fmt"

// two kinds of variable
// 1, package variable
// 2, local variable

// package variable prefer:
var a = 13
var b = int32(17)
var f = float32(3.14)

var (
	a1 = 13
	b1 = int32(17)
	f1 = float32(3.14)
)

// not init
var a2 int32
var f2 float64

func main() {
	var a int = 10
	fmt.Println(a)
	var b int
	fmt.Println(b)

	var (
		c int    = 128
		d int8   = 6
		s string = "hello"
		r rune   = 'A'
		t bool   = true
	)
	fmt.Println(c, d, s, r, t)

	var e, f, g int = 5, 6, 7
	fmt.Println(e, f, g)

	var h = 11
	fmt.Println(h)

	var i = int32(13)
	fmt.Println(i)

	// var j // syntax error: unexpected newline, expected type
	// fmt.Println(j)

	// short declaration
	a1 := 12
	b1 := 'A'
	c1 := "hello"
	fmt.Println(a1, b1, c1)

	a2, b2, c2 := 12, 'A', "hello"
	fmt.Println(a2, b2, c2)

	// local variable declaration prefer:
	a3 := int32(17)
	f3 := float32(3.14)
	s3 := []byte("hello, gopher!")
	fmt.Println(a3, f3, s3)

}

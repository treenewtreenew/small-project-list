package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	//int_length()
	//int_overflow()
	//number_literal()
	//new_number_literal()
	//split_number()
	//format_int()
	//float32_2_int()
	//float32_confusion()
	//exponent_float()
	complex_number()
}

// integer's length
func int_length() {
	var a, b = int(5), uint(6)
	var p uintptr = 0x12345678
	fmt.Println("signed integer a's length is ", unsafe.Sizeof(a))
	fmt.Println("unsigned integer b's length is ", unsafe.Sizeof(b))
	fmt.Println("uintprt's length is ", unsafe.Sizeof(p))
}

// integer overflow
func int_overflow() {
	var a int8 = 127
	a += 1
	fmt.Println(a)

	var ua uint8 = 1
	ua -= 2
	fmt.Println(ua)
}

// number literal
func number_literal() {
	a := 53
	b := 0700
	c1 := 0xaabbcc
	c2 := 0xddeeff
	fmt.Println(a, b, c1, c2)
}

func new_number_literal() {
	d1 := 0b10001111
	d2 := 0b10001111
	d3 := 0o700
	d4 := 0o700
	fmt.Println(d1, d2, d3, d4)
}

func split_number() {
	a := 5_3_7
	b := 0b_1000_0111
	c := 0_700
	d := 0o_700
	e := 0x_5c_6_d
	fmt.Println(a, b, c, d, e)
}

func format_int() {
	var a int8 = 59
	fmt.Printf("%b\n", a)
	fmt.Printf("%d\n", a)
	fmt.Printf("%o\n", a)
	fmt.Printf("%O\n", a)
	fmt.Printf("%x\n", a)
	fmt.Printf("%X\n", a)
}

func float32_2_int() {
	var f float32 = 139.8125
	bits := math.Float32bits(f)
	fmt.Printf("%b\n", bits)
}

func float32_confusion() {
	var f1 float32 = 16777216.0
	var f2 float32 = 16777217.
	fmt.Println(f1 == f2)
	fmt.Printf("f1 is %b\n", math.Float32bits(f1))
	fmt.Printf("f2 is %b\n", math.Float32bits(f2))
}

func exponent_float() {
	a := 6674.28e-2
	b := .12345e+5
	fmt.Println(a, b)

	c := 0x2.p10
	d := 0x1.Fp0
	fmt.Println(c, d)

	var f float64 = 123.45678
	fmt.Println(f)
	fmt.Printf("%e\n", f)
	fmt.Printf("%x\n", f)
}

func complex_number() {
	var c = complex(5, 6)
	r := real(c)
	i := imag(c)
	fmt.Println(c, r, i)
}

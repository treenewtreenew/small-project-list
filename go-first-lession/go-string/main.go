package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
	"unsafe"
)

// concurrent safety
// remove \0
// get length of string is quick, Constant time complexity
func string_cannot_be_changed() {
	var s string = "hello"
	// s[0] = 'k' // string in go can not be changed
	s = "gopher"

	s = `,_---~~~~~----._
	_,,_,*^____ _____*g*\"*,--,
	/ __/ /' ^. / \ ^@q f
	[ @f | @)) | | @)) l 0 _/
	\/ \~____ / __ \_____/ \
	| _l__l_ I
	} [______] I
	] | | | |
	] ~ ~ |
	| |
	| |`
	fmt.Println(s)
}

func string_content() {
	var s string = "中国人"
	fmt.Printf("the length of s = %d\n", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("0x%x", s[i])
	}
	fmt.Printf("\n")

	fmt.Println("the character count in s is ", utf8.RuneCountInString(s))
	for _, c := range s {
		fmt.Printf("0x%x", c)
	}
	fmt.Printf("\n")
}

func rune_in_go() {
	var a rune = 'a'
	var b rune = '中'
	var c rune = '\n'
	var d rune = '\''
	fmt.Println(string(a), string(b), string(c), string(d))

	a1 := '\u4e2d'
	b1 := '\U00004e2d'
	c1 := '\u0027'
	fmt.Println(string(a1), string(b1), string(c1))
}

func string_literal() {
	a := "abc\n"
	b := "中国人"
	c := "\u4e2d\u56fd\u4eba"
	d := "\U00004e2d\U000056fd\U00004eba"
	e := "中\u56fd\u4eba"
	f := "\xe4\xb8\xad\xe5\x9b\xbd\xe4\xba\xba"
	fmt.Println(a, b, c, d, e, f)
}

// rune -> []byte
func encodeRune() {
	var r rune = 0x4E2D
	fmt.Printf("the unicode character is %c\n", r)
	buf := make([]byte, 3)
	_ = utf8.EncodeRune(buf, r)
	fmt.Printf("utf8 representation is 0x%X\n", buf)
}

// []byte -> rune
func decodeRune() {
	var buf = []byte{0xE4, 0xB8, 0xAD}
	r, _ := utf8.DecodeRune(buf)
	fmt.Printf("the unicode character after decoding [0xE4, 0xB8, 0xAD] is %s\n", string(r))
}

func dumpBytesArray(arr []byte) {
	fmt.Printf("[")
	for _, b := range arr {
		fmt.Printf("%c ", b)
	}
	fmt.Printf("]\n")
}

func main() {
	var s = "hello"
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s)) // 将string类型变量地址显式转型为reflect.StringHeader
	fmt.Printf("0x%x\n", hdr.Data)                     // 0x10a30e0
	p := (*[5]byte)(unsafe.Pointer(hdr.Data))          // 获取Data字段所指向的数组的指针
	dumpBytesArray((*p)[:])                            // [h e l l o ]   // 输出底层数组的内容

	var ss = "中国人"
	fmt.Printf("0x%x\n", ss[0]) // 0xe4：字符“中” utf-8编码的第一个字节

	for i := 0; i < len(ss); i++ {
		fmt.Printf("index: %d, value: 0x%x\n", i, ss[i])
	}
	for i, v := range ss {
		fmt.Printf("index: %d, value: 0x%x\n", i, v)
	}
}

package main

import (
	"bytes"
	"fmt"
	"sync"
	"unsafe"

	book "GO-17/struct"
)

func main() {
	typeDefinition_1()
	withTheSameUnderlyingType()
	typeAlias()
	typeLiteral()
	typeBlock()
	typeStruct()
	usingStruct()
	emptyStruct()
	usingOtherStruct()
	selfContainUsingSliceMap()
	structDeclarationAndInit()
	zeroAvailable()
	usingLiteralInitStruct()
	structMemMap()
}

func typeDefinition_1() {
	type T1 string // Underlying Type is string
	var s1 T1
	fmt.Printf("my type is %T\n", s1)

	type T2 T1 // Underlying Type is also string
	var s2 T2
	fmt.Printf("my type is %T\n", s2)
}

func withTheSameUnderlyingType() {
	type T1 int
	type T2 T1
	type T3 string

	var n1 T1
	var n2 T2 = 5
	n1 = T1(n2)
	var s T3 = "hello"
	// n1 = T1(s) // cannot convert s (variable of type T3) to type
	fmt.Println(n1)
	fmt.Println(s)
}

func typeLiteral() {
	type M map[string]int
	type StrArr []string
}

func typeBlock() {
	type (
		T1 int
		T2 T1
		T3 string
	)
	var t1, t4 T1
	var t2 T2
	var t3 T3
	fmt.Printf("%T %T %T %T", t1, t2, t3, t4)
}

func typeAlias() {
	type T = string
	var s string = "hello"
	var t T = s
	fmt.Printf("%T\n", t)
}

func typeStruct() {
	type (
		T1 int
		T2 T1
		T3 string
	)

	type T struct {
		field1 T1
		field2 T2
		field3 T3
	}
}

func usingStruct() {
	var b book.Book
	b.Title = "The Go Programming Language"
	b.Pages = 800
	fmt.Println(b)
}

func emptyStruct() {
	type Empty struct{}
	//var s Empty
	// println(unsafe.Sizeof(s))

	// var c = make(chan Empty) // 声明一个元素类型为Empty的channel
	// c <- Empty{}             // 向channel写入一个“事件”
}

func usingOtherStruct() {
	type Person struct {
		Name  string
		Phone string
		Addr  string
	}

	type Book struct {
		Title  string
		Author Person
	}

	var book Book
	println(book.Author.Phone)
}

func selfContain() {
	// type T struct {
	// 	t T // invalid recursive type: T refers to itself
	// }
}

func recursiveContain() {
	// type T1 struct {
	// 	t T2
	// }
	// type T2 struct {
	// 	t T1
	// }
}

func selfContainUsingSliceMap() {
	type T struct {
		t  *T           // ok
		st []T          // ok
		m  map[string]T // ok

	}
	var t1 T
	fmt.Println(t1)
	fmt.Printf("%T", t1)
}

func structDeclarationAndInit() {
	type Person struct {
		Name  string
		Phone string
		Addr  string
	}

	type Book struct {
		Title  string
		Author Person
	}

	var book1 Book
	var book2 = Book{}
	book3 := Book{}
	fmt.Printf("%T %T %T\n", book1, book2, book3)
	fmt.Println(unsafe.Sizeof(book1), unsafe.Sizeof(book2), unsafe.Sizeof(book3))
}

func zeroAvailable() {
	var mu sync.Mutex
	mu.Lock()
	mu.Unlock()

	var b bytes.Buffer
	b.Write([]byte("Hello, Go"))
	fmt.Println(b.String())
}

func usingLiteralInitStruct() {
	type Book struct {
		Title   string
		Pages   int
		Indexes map[string]int
	}
	var book Book = Book{"The Go Programming language", 800, make(map[string]int)}
	fmt.Println(book)

	// init struct using field:value style
	book = Book{
		Title:   "The go programming language",
		Indexes: make(map[string]int),
		Pages:   100,
	}
	fmt.Println(book)
}

func structMemMap() {
	type T1 struct {
		b byte
		i int64
		u uint16
	}
	var t1 T1
	fmt.Println(unsafe.Sizeof(t1))

	type T2 struct {
		b byte
		u uint16
		i int64
	}
	var t2 T2
	fmt.Println(unsafe.Sizeof(t2))

	type T3 struct {
		i int64
		b byte
		u uint16
	}
	var t3 T3
	fmt.Println(unsafe.Sizeof(t3))
}

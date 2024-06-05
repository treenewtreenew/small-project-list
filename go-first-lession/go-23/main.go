package main

import "fmt"

// robust 3 factors
// 1, don't trust any external input
// 2, don't ignore any error
// 3, don't assume the exception would not occur

// exceptions are not same as errors. Errors can be predicted and
// we could create some error processing code snippets.

// While, the exceptions could not be predicted in advance.
// hardware exception, os exception, go runtime exception, exceptions caused by
// unrevealed bugs.

// panicking
func foo() {
	println("call foo")
	bar()
	println("exit foo")
}

func bar() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("recover the panic: ", e)
		} else {
			fmt.Println("There is no panic")
		}
	}()
	println("call bar")
	panic("panic occurs in bar")
	zoo()
	println("exit bar")
}

func zoo() {
	println("call zoo")
	println("exit zoo")
}

func callFoo() {
	println("call callFoo")
	foo()
	println("exit callFoo")
}

func main() {
	callFoo()
}

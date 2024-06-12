package main

import (
	"context"
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"sync"
)

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
	print("exit foo")
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

// go http server example
// $GOROOT/src/net/http/server.go
// serve a new connection
func (c *conn) serve(ctx context.Context) {
	c.remoteAddr = c.rwc.RemoteAddr().String()
	ctx = context.WithValue(ctx, LocalAddrContextKey, c.rwc.LocalAddr())
	defer func() {
		if err := recover(); err != nil && err != ErrAbortHandler {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			c.server.logf("http: panic serving %v: %v\n%s", c.remoteAddr, err, buf)
		}
		if !c.hijacked() {
			c.close()
			c.setState(c.rwc, StateClosed, runHooks)
		}
	}()

}

// 在 Go 标准库中，大多数 panic 的使用都是充当类似断言的作用的。
// $GOROOT/src/encoding/json/encode.go
func (w *reflectWithString) resolve() error {
	// ... ...
	switch w.k.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		w.ks = strconv.FormatInt(w.k.Int(), 10)
		return nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		w.ks = strconv.FormatUint(w.k.Uint(), 10)
		return nil
	}
	panic("unexpected map key type")
}

// using defer to simplify source
func doSomething1() error {
	var mu sync.Mutex
	mu.Lock()

	r1, err := OpenResource1()
	if err != nil {
		mu.Unlock()
		return err
	}

	r2, err := OpenResource2()
	if err != nil {
		r1.Close()
		mu.Unlock()
		return err
	}

	r3, err := OpenResource3()
	if err != nil {
		r2.Close()
		r1.Close()
		mu.Unlock()
		return err
	}

	// 使用r1，r2, r3
	err = doWithResources()
	if err != nil {
		r3.Close()
		r2.Close()
		r1.Close()
		mu.Unlock()
		return err
	}

	r3.Close()
	r2.Close()
	r1.Close()
	mu.Unlock()
	return nil
}

// 在 Go 中，只有在函数（和方法）内部才能使用 defer；
// defer 关键字后面只能接函数（或方法），这些函数被称为 deferred 函数。
// defer 将它们注册到其所在 Goroutine 中，用于存放 deferred 函数的栈数
// 据结构中，这些 deferred 函数将在执行 defer 的函数退出前，按后进先出
// （LIFO）的顺序被程序调度执行（如下图所示）。
func doSomething2() error {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()

	r1, err := OpenResource1()
	if err != nil {
		return err
	}
	defer r1.Close()

	r2, err := OpenResource2()
	if err != nil {
		return err
	}
	defer r2.Close()

	r3, err := OpenResource3()
	if err != nil {
		return err
	}
	defer r3.Close()

	// 使用r1，r2, r3
	return doWithResources()
}

// defer 使用的几个注意事项
// 1, 对于有返回值的自定义函数或方法，返回值会在 deferred 函数被调度执行的时候被自动丢弃。
// go builtin function
/*
Functions:
  append cap close complex copy delete imag len
  make new panic print println real recover
*/
// deferred 函数的参数值都是在注册的时候进行求值的。

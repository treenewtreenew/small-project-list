package main

import (
	"errors"
	"fmt"
	"io"
	"net"
)

// if return 0 then success
// simple

// c
// func Fprintf(w io.Writer, format string, a...interface{}) (n int, err error)

type error interface {
	Error() string
}

func exampleError() {
	err := errors.New("your first demo error")
	println(err.Error())
	i := 6
	errWithCtx := fmt.Errorf("index %d is out of bounds", i)
	println(errWithCtx.Error())
}

// $GOROOT/src/net/net.go
type OpError struct {
	Op  string
	Net string
	// Source Addr
	// addr Addr
	Err error
}

func isCommonNetReadError(err error) bool {
	if err == io.EOF {
		return true
	}
	if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
		return true
	}
	if oe, ok := err.(*net.OpError); ok && oe.Op == "read" {
		return true
	}
	return false
}

// wrapped error
func wrappedError() {
	var errSentinel = errors.New("the underlying sentinel error")
	err1 := fmt.Errorf("wrap sentinel: %w", errSentinel)
	err2 := fmt.Errorf("wrap sentinel %w", err1)
	println(err2 == errSentinel)
	if errors.Is(err2, errSentinel) {
		println("err2 is errSentinel")
	}
	if errors.Is(err1, errSentinel) {
		println("err1 is errSentinel")
	}
	if errors.Is(err2, err1) {
		println("err2 is err1")
		return
	}
	println()
}

// wrapped error : errors.As
type MyError struct {
	e string
}

func (e *MyError) Error() string {
	return e.e
}

func checkWrappedErrorUsingAs() {
	var err = &MyError{"MyError error demo"}
	err1 := fmt.Errorf("wrap err: %w", err)
	err2 := fmt.Errorf("wrap err1: %w", err1)
	var e *MyError
	if errors.As(err2, &e) {
		println("MyError is on the chain of err2")
		// errors.As 会将匹配到的错误值存储到 As 函数的第二个参数中，
		// 这也是为什么 println(e == err)输出 true 的原因。
		println(e == err)
		return
	}
	println("MyError is not on the chain of err2")
}

func main() {
	exampleError()
	wrappedError()
	checkWrappedErrorUsingAs()
}

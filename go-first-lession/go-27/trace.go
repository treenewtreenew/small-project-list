package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

// 增加 Goroutine 标识
var goroutineSpace = []byte("goroutine ")

func curGoroutineID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	// Parse the 4707 out of "goroutine 4707 ["
	b = bytes.TrimPrefix(b, goroutineSpace)
	i := bytes.IndexByte(b, ' ')
	if i < 0 {
		panic(fmt.Sprintf("No space found in %q", b))
	}
	b = b[:i]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse goroutine ID out of %q: %v", b, err))
	}
	return n
}

func Trace() func() {

	// 自动获取所跟踪函数的函数名
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("not found caller")
	}

	fn := runtime.FuncForPC(pc)
	name := fn.Name()

	gid := curGoroutineID()
	fmt.Printf("g[%05d]: enter: [%s]\n", gid, name)
	return func() {
		fmt.Printf("g[%05d]: enter: [%s]\n", gid, name)
	}
}

func foo() {
	defer Trace()()
	bar()
}

func bar() {
	defer Trace()()
}

// trace2/trace.go
func A1() {
	defer Trace()()
	B1()
}

func B1() {
	defer Trace()()
	C1()
}

func C1() {
	defer Trace()()
	D()
}

func D() {
	defer Trace()()
}

func A2() {
	defer Trace()()
	B2()
}
func B2() {
	defer Trace()()
	C2()
}
func C2() {
	defer Trace()()
	D()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		A2()
		wg.Done()
	}()

	A1()
	wg.Wait()
}

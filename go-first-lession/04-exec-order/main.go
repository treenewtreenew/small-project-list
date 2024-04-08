package main

import (
	"fmt"

	_ "github.com/bigwhite/prog-init-order/pkg1"
	_ "github.com/bigwhite/prog-init-order/pkg2"
)

func init() {
	fmt.Println("init invoked")
}

// executible go program must have a main func.
// main.main is the entrance func, while it maybe
// not the first func to be executed. Because in go, there are init func.
// there could be multiple init func

// init order: const, variable, init func

// init func using scenarios:
// 1, reset package variable
func main() {
	// init() // calling func init manually is compiling error
}

var (
	_  = constInitCheck()
	v1 = variableInit("v1")
	v2 = variableInit("v2")
)

const (
	c1 = "c1"
	c2 = "c2"
)

func constInitCheck() string {
	if c1 != "" {
		fmt.Println("main: const c1 has been initialized")
	}
	if c2 != "" {
		fmt.Println("main: const c2 has been initialized")
	}
	return ""
}

func variableInit(name string) string {
	fmt.Printf("main: var %s has been initialized\n", name)
	return name
}

func init() {
	fmt.Println("main: first init func invoked")
}

func init() {
	fmt.Println("main: second init func invoked")
}

package main

import "fmt"

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

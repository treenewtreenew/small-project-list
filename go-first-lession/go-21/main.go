package main

import (
	"fmt"
	"net/http"
)

// function literal
var foo = func() {
	println("this is a func definition")
}

// variable parameters declaration
func vFunc(i int, s string, a ...interface{}) (n int, err error) {
	println(i)
	println(s)
	for i, v := range a {
		println(i)
		println(v)
	}
	return 1, nil
}

func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, Gopher!\n")
}

func times(x int, y int) int {
	return x * y
}

func partialTimes(x int) func(int) int {
	return func(y int) int {
		return times(x, y)
	}
}

func main() {
	foo()
	vFunc(1, "1", "a", "b", 2)
	// cannot use greeting (value of type func(w http.ResponseWriter, r *http.Request))
	// as http.Handler value in argument to http.ListenAndServe:
	// func(w http.ResponseWriter, r *http.Request) does not implement
	// http.Handler (missing method ServeHTTP)
	// http.ListenAndServe(":8080", greeting)
	// below is correct
	// http.ListenAndServe(":8080", http.HandlerFunc(greeting))
	times2 := partialTimes(2)
	println(times2(3))
	time3 := partialTimes(3)
	println(time3(3))
}

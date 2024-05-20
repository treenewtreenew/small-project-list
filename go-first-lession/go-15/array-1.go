package main

// var arr [N]T

func foo(arr [5]int) {}

func funcparameter() {
	var arr1 [5]int
	var arr2 [6]int
	var arr3 [5]string

	foo(arr1)
	foo(arr2) // parameter type [5]int, arr2 is [6]int
	foo(arr3) // parameter type [6]int, arr3 is [5]string
}

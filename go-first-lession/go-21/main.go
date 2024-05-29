package main

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

func main() {
	foo()
	vFunc(1, "1", "a", "b", 2)
}

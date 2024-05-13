package main

import "runtime"

func main() {
	printFirstIf()
	recursiveIf()
}

func printFirstIf() {
	if runtime.GOOS == "Linux" {
		println("we are on linux os")
	}
	if runtime.GOOS == "windows" {
		println("we are on windows os", runtime.GOOS)
	}

	if (runtime.GOOS != "linux") && (runtime.GOARCH == "amd64") &&
		(runtime.Compiler != "gccgo") {
		println("we are using standard go compiler on windows os for amd64")
	}
}

func recursiveIf() {
	if a, c := -1, 2; a > 0 {
		println(a)
	} else if b := 1; b > 0 {
		println(a, b)
	} else {
		println(a, b, c)
	}
}

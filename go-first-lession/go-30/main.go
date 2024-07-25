package main

import (
	"io"
	"log"
	"os"
	"strings"
)

type MyReadWriter interface {
	io.Reader
	io.Writer
}

type MyReader struct {
	io.Reader
	N int64
}

type MyReader2 struct {
	MyReader
	MyReadWriter
}

func main() {
	r := strings.NewReader("hello, gopher!\n")
	lr := io.LimitReader(r, 4)
	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}
}

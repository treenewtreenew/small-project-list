package main

import (
	"bytes"
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

	r1 := strings.NewReader("hello, gopher!\n")
	r2 := CapReader(io.LimitReader(r1, 4))
	if _, err := io.Copy(os.Stdout, r2); err != nil {
		log.Fatal(err)
	}
}

func CapReader(r io.Reader) io.Reader {
	return &capitalizedReader{r: r}
}

type capitalizedReader struct {
	r io.Reader
}

func (r *capitalizedReader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	if err != nil {
		return 0, err
	}

	q := bytes.ToUpper(p)
	for i, v := range q {
		p[i] = v
	}
	return n, err
}

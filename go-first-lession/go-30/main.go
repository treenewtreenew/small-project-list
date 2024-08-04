package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
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

func validateAuth(s string) error {
	if s != "123456" {
		return fmt.Errorf("%s", "bad auth token")
	}
	return nil
}

func greetings(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func logHandler(h http.Handler)  http.Handler{
	return  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)){
		t := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t)
		h.ServeHTTP(w, r)
	}
}

func authHandler(h http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r * http.Request){
		err := validateAuth(r.Header.Get("auth"))
		if err != nil{
			http.Error(w, "bad auth param", http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	})
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

	http.ListenAndServe(":8080", logHandler(authHandler(http.HandlerFunc(greetings))))
}

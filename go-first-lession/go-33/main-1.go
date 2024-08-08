package main

import (
	"fmt"
	"time"
)

type signal struct{}

func worker_1() {
	println("worker is working...")
	time.Sleep(5 * time.Second)
}

func spawn(f func()) <-chan signal {
	c := make(chan signal)
	go func() {
		println("worker start to work...")
		f()
		c <- signal{}
	}()
	return c
}

func main_1() {
	println("start a worker...")
	c := spawn(worker_1)
	<-c
	fmt.Println("worker work done!")
}

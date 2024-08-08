package main

import (
	"sync"
	"time"
)

// communicating sequential process is composed of 2 parts
// 1st part: goroutine
// 2nd part: channel

func correctChanSendAndReceiveWithBuffer() {
	ch1 := make(chan int, 5)
	ch1 <- 13
	n := <-ch1
	println(n)
}

// fatal error: all goroutines are asleep - deadlock!
func errorChanSendAndReceiveWithoutBuffer() {
	ch1 := make(chan int)
	ch1 <- 13
	n := <-ch1
	println(n)
}

// correct: 将接收操作，或者发送操作放到另外一个 Goroutine 中
func correctChanSendAndReceiveWithoutBuffer() {
	ch1 := make(chan int)
	go func() {
		ch1 <- 13
	}()
	n := <-ch1
	println(n)
}

// buffer is full
func bufferIsFull() {
	ch2 := make(chan int, 1)
	n := <-ch2

	ch3 := make(chan int, 1)
	ch3 <- 17
	ch3 <- -27

	println(n)
}

// syntax
func sendOnlyAndReadOnly() {
	ch1 := make(chan<- int, 1)
	ch2 := make(<-chan int, 1)
	// <-ch1  invalid operation: cannot receive from send-only channel ch1 (variable of type chan<- int)
	// ch2 <- 13 invalid operation: cannot send to receive-only channel ch2 (variable of type <-chan int)
	ch1 <- 13
	n := <-ch2
	println(n)
}

// produce
func produce(ch chan<- int) {
	for i := 10; i < 10; i++ {
		ch <- i + 1
		time.Sleep(time.Second)
	}
	close(ch)
}

// consume
func consume(ch <-chan int) {
	for n := range ch {
		println(n)
	}
}

// send to a closed channel
func sendToClosed() {
	ch := make(chan int, 5)
	close(ch)
	ch <- 13 // panic: send on closed channel
}

func main_0() {

	ch := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		produce(ch)
		wg.Done()
	}()

	go func() {
		consume(ch)
		wg.Done()
	}()

	wg.Wait()
}

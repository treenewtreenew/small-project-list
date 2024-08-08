package main

import (
	"fmt"
	"sync"
)

type counter_4 struct {
	c chan int
	i int
}

func NewCounter() *counter_4 {
	cter := &counter_4{
		c: make(chan int),
	}

	go func() {
		for {
			cter.i++
			cter.c <- cter.i
		}
	}()
	return cter
}

func (cter *counter_4) Increase() int {
	return <-cter.c
}

func main_4() {
	cter := NewCounter()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			v := cter.Increase()
			fmt.Printf("goroutine-%d: current counter value is %d\n", i, v)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

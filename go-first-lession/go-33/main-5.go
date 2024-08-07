package main

import (
	"log"
	"sync"
	"time"
)

// 第二种用法：用作计数信号量（counting semaphore）
var active = make(chan struct{}, 3)
var jobs = make(chan int, 10)

func main_5() {
	go func() {
		for i := 0; i < 8; i++ {
			jobs <- (i + 1)
		}
		close(jobs)
	}()

	var wg sync.WaitGroup

	for j := range jobs {
		wg.Add(1)
		go func(j int) {
			active <- struct{}{}
			log.Printf("handle job: %d\n", j)
			time.Sleep(2 * time.Second)
			<-active
			wg.Done()
		}(j)
	}
	wg.Wait()
}

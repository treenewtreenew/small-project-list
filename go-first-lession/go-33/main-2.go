package main

import (
	"fmt"
	"sync"
	"time"
)

func worker_2(i int) {
	fmt.Printf("worker %d: is working...\n", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("worker %d: works done\n", i)
}

type signal_2 struct{}

func spawnGroup(f func(i int), num int, groupSignal <-chan signal_2) <-chan signal_2 {
	c := make(chan signal_2)
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			<-groupSignal
			fmt.Printf("worker %d: start to work...\n", i)
			f(i)
			wg.Done()
		}(i + 1)
	}

	go func() {
		wg.Wait()
		c <- signal_2{}
	}()
	return c
}

func main_2() {
	fmt.Println("start a group of workers...")
	groupSignal := make(chan signal_2)
	c := spawnGroup(worker_2, 5, groupSignal)
	time.Sleep(5 * time.Second)
	fmt.Println("the group of workers start to work...")
	close(groupSignal)
	<-c
	fmt.Println("the group of workers work done!")
}

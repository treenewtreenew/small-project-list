package main

import (
	"fmt"
	"sync"
	"time"
)

// len(channel) 的应用
// 将“判空与读取”放在一个“事务”中，将“判满与写入”放在一个“事务”中

func producer(c chan<- int) {
	var i int = 1
	for {
		time.Sleep(2 * time.Second)
		ok := trySend(c, i)
		if ok {
			fmt.Printf("[producer]: send [%d] to channel\n", i)
			i++
			continue
		}
		fmt.Printf("[producer]: try send [%d], but channel is full\n", i)
	}
}

func tryRecv(c <-chan int) (int, bool) {
	select {
	case i := <-c:
		return i, true
	default:
		return 0, false
	}
}

func trySend(c chan<- int, i int) bool {
	select {
	case c <- i:
		return true
	default:
		return false
	}
}

func consumer(c <-chan int) {
	for {
		i, ok := tryRecv(c)
		if !ok {
			fmt.Println("[consumer]: try to recv from channel, but the channel is empty")
			time.Sleep(1 * time.Second)
			continue
		}
		fmt.Printf("[consumer]: recv [%d] from channel\n", i)
		if i >= 3 {
			fmt.Println("[consumer]: exit")
			return
		}
	}
}

func main_6() {
	var wg sync.WaitGroup
	c := make(chan int, 3)
	wg.Add(2)

	go func() {
		consumer(c)
		wg.Done()
	}()

	go func() {
		producer(c)
		wg.Done()
	}()

	wg.Wait()
}

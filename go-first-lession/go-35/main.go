package main

import (
	"errors"
	"sync"
)

// pool 的创建与销毁
// pool 中 worker（Goroutine）的管理
// task 的提交与调度

type Pool struct {
	capacity int 			// workerpool size
	active chan struct{} 	// active channel
	tasks chan Task			// task channel
	wg sync.WaitGroup		// waiting all workers to finish and quit
	quit chan struct
}

func New(capacity int) *Pool {
	if capacity <= 0 {
		capacity = defaultCapacity
	}
	if capacity > maxCapacity {
		capacity = maxCapacity
	}

	p := &Pool{
		capacity: capacity,
		tasks: make(chan Task)
		quit: make(chan struct{})
		active: make(chan struct{}, capacity),
	}

	fmt.Printf("workerpool start\n")
	
	go p.run()
	return p
}

func (p *Pool) run() {
	idx := 0

	for {
		select {
		case <-p.quit:
			return
		case p.active <- struct{}{}:
			idx++
			p.newWorker(idx)
		}
	}
}

func (p *Pool) newWorker(i int) {
	p.wg.Add(1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("worker[%03d]: recover panic[%s] and exit\n", i, err)
				<-p.active
			}
			p.wg.Done()
		}()

		fmt.Printf("worker[%03d]: start \n", i)

		for {
			select {
			case <-p.quit:
				fmt.Printf("worker[%03d]: exit\n", i)
				<-p.active
				return
			case t := <-p.tasks:
				fmt.Printf("worker[%03d]: receive a task\n", i)
				t()
			}
		}
	}()
}

var ErrWorkerPoolFreed = errors.New("workerpool freed")

func (p *Pool) Schedule(t task) error {
	select {
	case <-p.quit:
		return ErrWorkerPoolFreed
	case p.tasks <- t:
		return nil
	}
}
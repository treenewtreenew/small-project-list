package main

// Goroutine Scheduler
// Go runtime
// G-M model, G-P-M model

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int, 5)
	println(ch1, ch2)
}

package main

func main() {
	var c chan int
	<-c    //阻塞
	c <- 1 //阻塞
}

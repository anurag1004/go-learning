package main

import "fmt"

func main() {
	/*
		The optional <- operator specifies the channel direction i.e send or receive
		if no direction is given, the channel is bi-directional
		<-chan int (channel which can only receive ints) RECEIVE ONLY CHANNEL
		chan<-int (channel in which we can send ints) SEND ONLY CHANNEL
	*/
	recvChan1 := producer()
	sumChan := consumer(recvChan1)
	for sum := range sumChan {
		fmt.Println(sum)
	}
}
func producer() <-chan int {
	out := make(chan int)
	go func() {
		var i int = 1
		for i <= 10 {
			out <- i
			i++
		}
		close(out)
	}()
	return out
}
func consumer(ch <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum int = 0
		for num := range ch {
			sum += num
		}
		out <- sum
		close(out)
	}()
	return out
}

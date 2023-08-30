package main

import "fmt"

func main() {
	for sqrs := range emitSquares(emitSquares(streamInts(1, 2, 3, 4, 5, 6, 7, 8))) {
		fmt.Println(sqrs)
	}
}

func streamInts(nums ...int) <-chan int {
	// the job of this pipeline is to stream each ints to channel
	out := make(chan int)
	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()
	return out // inbound channel for next pipeline
}

func emitSquares(inboundChan <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range inboundChan {
			out <- num * num
		}
		close(out)
	}()
	return out
}

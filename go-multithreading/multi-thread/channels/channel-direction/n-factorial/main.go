package main

import "fmt"

func main() {
	/*
		prodChan := producer(5)
		consumeChan := consumer(prodChan)
		for fact := range consumeChan {
			fmt.Println(fact)
		}
	*/
	for fact := range consumer(producer(5)) {
		fmt.Println(fact)
	}
}
func producer(N int) <-chan int {
	out := make(chan int)
	go func() {
		for N > 0 {
			out <- N
			N--
		}
		close(out)
	}()
	return out
}
func consumer(ch <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var fact int = 1
		for num := range ch {
			fact *= num
		}
		out <- fact
		close(out)
	}()
	return out
}

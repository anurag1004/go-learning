package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int)

	wg.Add(2)
	go foo(ch)
	go bar(ch)
	go func() {
		wg.Wait()
		close(ch) // if we dont close the channel here
		// the receving end will always listen for data
		// and our main thread will be blocked
	}()

	for data := range ch { // this loop will terminate only when channel receives a closing signal
		// otherwise our program will simply stuck here indefinetly
		fmt.Println(data)
	}
	fmt.Println("DONE")
}
func foo(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i // 1-5
	}
	wg.Done()
}
func bar(ch chan int) {
	for i := 5; i <= 10; i++ {
		ch <- i // 5-10
	}
	wg.Done()
}

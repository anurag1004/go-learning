package main

import (
	"fmt"
	"time"
)

func main() {
	N := 3
	ch := make(chan int, N) // buffer size of 5, act like semaphore
	for i := 0; i < 5; i++ {
		go foo(ch, i)
	}
	time.Sleep(10 * time.Second)
}
func foo(ch chan int, id int) { // max N goroutines can enter
	fmt.Printf("FOO: TID:%d Waiting...[%s]\n", id, time.Now().Format("2006-01-02 15:04:05"))
	ch <- 1 // send operation will be blocked if buffer is full
	fmt.Printf("FOO: TID:%d Inside [%s]\n", id, time.Now().Format("2006-01-02 15:04:05"))
	bar(id)
	<-ch //
	fmt.Printf("FOO: TID:%d DONE...[%s]\n", id, time.Now().Format("2006-01-02 15:04:05"))
}
func bar(id int) {
	// some task
	time.Sleep(1 * time.Second)
	fmt.Printf("BAR: TID:%d, square:%d\n", id, id*id)
}

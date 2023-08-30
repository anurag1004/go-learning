package main

import (
	"fmt"
	"time"
)

func main() {
	N := 4
	ch := make(chan int, N)
	for i := 0; i < N+4; i++ {
		go foo(i, ch)
	}
	time.Sleep(10 * time.Second)
}
func foo(id int, ch chan int) {
	fmt.Printf("TID: %d Waiting...[%s]\n", id, time.Now().Format("2006-01-02 15:04:05"))
	ch <- 1
	fmt.Printf("TID: %d Inside...[%s]\n", id, time.Now().Format("2006-01-02 15:04:05"))
	time.Sleep(1 * time.Second)
	<-ch
	fmt.Printf("TID: %d Done...[%s]\n", id, time.Now().Format("2006-01-02 15:04:05"))
}

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int // this will be out shared data
func init() {
	runtime.GOMAXPROCS(1) // simulate unicore
}
func main() {
	wg.Add(2)
	// these two go routine will now run concurrently
	go someFunc(1)
	go someFunc(2)
	wg.Wait()
	fmt.Printf("Final counter:%d", counter)
}
func someFunc(id int) {
	for i := 0; i < 10; i++ {
		// critical section starts
		counter++ // critical section
		// critical section ends
		fmt.Printf("i:%d, TID:%d, counter: %d\n", i, id, counter)
		time.Sleep(500 * time.Millisecond)
	}
	wg.Done()
}

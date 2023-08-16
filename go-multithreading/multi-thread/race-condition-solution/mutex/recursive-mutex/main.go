package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.Mutex
var counter int // this will be out shared data
func main() {
	wg.Add(2)
	go someRecursiveFunc(1, 10) // tid1
	go someRecursiveFunc(2, 10) // tid2
	wg.Wait()
	fmt.Printf("Final Counter: %d\n", counter)
}

func someRecursiveFunc(id int, iter int) {
	if iter <= 0 {
		wg.Done()
		return
	}
	fmt.Printf("Entery section - TID: %d, iter: %d, waiting to enter CS...\n", id, iter)

	mutex.Lock()

	fmt.Printf("Mutex aquired by TID: %d\n", id)
	counter++
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("iter:%d, TID:%d, counter: %d\n", iter, id, counter)

	mutex.Unlock()

	someRecursiveFunc(id, iter-1)
	fmt.Printf("Mutex released by TID: %d, iter: %d\n", id, iter)
}
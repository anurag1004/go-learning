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
	go someFunc(1) // tid1
	go someFunc(2) // tid2
	// go someOtherFunc(1)
	// go someOtherFunc(2)
	wg.Wait()
	fmt.Printf("Final Counter: %d\n", counter)
}

func someFunc(id int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Entery section - Thread: %d, waiting to enter CS...\n", id)
		mutex.Lock() // aquire mutex
		fmt.Printf("Mutex aquired by: %d\n", id)
		// critical section starts
		counter++
		// critical section ends
		fmt.Printf("i:%d, TID:%d, counter: %d\n", i, id, counter)
		time.Sleep(500 * time.Millisecond)
		mutex.Unlock()
		fmt.Printf("Mutex released by: %d\n", id)
	}
	wg.Done()
}
func someOtherFunc(id int) {
	// here we synchronize the entire funtion
	mutex.Lock() // aquire mutex
	for i := 0; i < 10; i++ {
		// critical section starts
		counter++
		// critical section ends
		fmt.Printf("i:%d, TID:%d, counter: %d\n", i, id, counter)
		time.Sleep(500 * time.Millisecond)
	}
	mutex.Unlock()
	wg.Done()
}

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutexA sync.Mutex
var mutexB sync.Mutex

func main() {
	wg.Add(2)
	go foo(1)
	go foo(2)
	wg.Wait()
}
func foo(id int) {
	fmt.Printf("TID: %d, waiting to aquire mutexA [%s]\n", id, time.Now().Format("2006-01-02 15:04:05"))
	{
		mutexA.Lock()
		fmt.Printf("TID: %d, Acquired mutexA [%s]\n", id, time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(10 * time.Second)
		fmt.Printf("TID: %d, waiting to aquire mutexB [%s]\n", id, time.Now().Format("2006-01-02 15:04:05"))
		{
			mutexB.Lock()
			fmt.Printf("TID: %d, Acquired mutexB [%s]\n", id, time.Now().Format("2006-01-02 15:04:05"))
			time.Sleep(10 * time.Second)
			mutexB.Unlock()
		}
		fmt.Printf("MutexB released by TID:%d [%s]\n", id, time.Now().Format("2006-01-02 15:04:05"))
		mutexA.Unlock()
		fmt.Printf("MutexA released by TID:%d [%s]\n", id, time.Now().Format("2006-01-02 15:04:05"))
	}
	wg.Done()
}

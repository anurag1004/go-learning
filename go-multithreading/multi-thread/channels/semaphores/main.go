package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// raw semaphore
type Semaphore struct {
	val int32
	max int32
}

func (s *Semaphore) Signal() bool {
	if atomic.LoadInt32(&(*s).val) >= s.max {
		// semaphore max capacity reached
		return false
	}
	atomic.AddInt32(&(*s).val, 1)
	return true
}
func (s *Semaphore) Wait() {
	for {
		// wait signal
		if sp.waitHelper() {
			break
		}
	}
}
func (s *Semaphore) waitHelper() bool {
	if atomic.LoadInt32(&(*s).val) == 0 {
		// semaphore is full
		return false
	}
	atomic.AddInt32(&(*s).val, -1)
	return true
}

var sp = &Semaphore{
	val: 3,
	max: 3,
}
var wg sync.WaitGroup

func main() {
	wg.Add(4)
	go foo(1)
	go foo(2)
	go foo(3)
	go foo(4)
	wg.Wait()
}
func foo(id int) {
	fmt.Printf("Thread %d, waiting for semaphore [%s]\n", id, time.Now().Format("2006-01-02 15:04:05"))
	sp.Wait()
	fmt.Printf("Thread %d, acquired semaphore [%s]\n", id, time.Now().Format("2006-01-02 15:04:05"))
	time.Sleep(5 * time.Second)
	sp.Signal()
	fmt.Printf("Thread %d, released semaphore [%s]\n", id, time.Now().Format("2006-01-02 15:04:05"))
	wg.Done()
}

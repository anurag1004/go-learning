package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

/*
In Go, concurrency is achieved using goroutines and channels,
while parallelism is achieved using multiple processors or cores.
To determine whether your code is running concurrently or in parallel,
you can use the runtime.NumCPU() function to get the number of available CPUs on your system,
and the runtime.GOMAXPROCS() function to get the maximum number of CPUs that Go will use for parallel execution.
If the number of available CPUs is greater than the maximum number of CPUs that Go will use,
then your code is running concurrently but not in parallel.
If the number of available CPUs is equal to or less than the maximum number of CPUs that Go will use,
then your code is running concurrently and in parallel.
*/
func main() {
	// Three threads are running here..
	// main thread ... (parent thread which is spawning all these child threads)
	/*
		go doSomething()
		go doSomethingElse()
	*/
	// this code would...simple run and will not show any output
	// because the main thread will simply spawn two go routines and will finish its excution
	fmt.Printf("Number of CPUs: %d\n", runtime.NumCPU())
	fmt.Printf("Maximum number of CPUs: %d\n", runtime.GOMAXPROCS(runtime.NumCPU())) // be default go will use as many cpus which are available
	// add wait groups
	wg.Add(2)            // we are directing the main thread to wait for any two go routine to finish
	go doSomething()     // then this
	go doSomethingElse() // last one to complete, i.e we'll not be able to see the whole output for this go routine
	go doSomethingMore() // this will finish first
	// we need to tell it to wait for its child threads to complete its job
	wg.Wait()
}
func doSomething() {
	for i := 0; i < 10; i++ {
		fmt.Printf("doSomething: %d\n", i)
		time.Sleep(500 * time.Millisecond) // add this because we want to simulate some heavy operation here
		// if we simply run small loops then we'll see sequential excution because of the speed at which it excutes
	}
	wg.Done()
}
func doSomethingElse() {
	for i := 0; i < 20; i++ {
		fmt.Printf("doSomethingElse: %d\n", i)
		time.Sleep(500 * time.Millisecond)
	}
	wg.Done()
}
func doSomethingMore() {
	for i := 0; i < 10; i++ {
		fmt.Printf("doSomethingMore: %d\n", i)
		time.Sleep(100 * time.Millisecond)
	}
	wg.Done()
}

package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

/*
Fan-out/fan-in refers to the pattern of executing multiple functions concurrently
and then performing some aggregation on the results
*/
var N int
var nums []int
var numProcessors int
var fakeProcessTime time.Duration

func init() {
	log.Println("init...")
	N = 100000
	numProcessors = 1000
	fakeProcessTime = 100 * time.Millisecond
	nums = make([]int, N)
	fillNums(nums, N)

	expectedTime := (float64(N/numProcessors) * float64(fakeProcessTime.Milliseconds())) / 1000
	fmt.Printf("Nums:%d,\n"+
		"numProcessors:%d,\n"+
		"FakeProcessTime:%s,\n"+
		"Expected time:%0.3fs\n",
		N, numProcessors, fakeProcessTime, expectedTime,
	)
}
func main() {
	done := make(chan struct{})
	inputChan := genNums(done, nums...) // input stream
	// spawn numProcessors
	processors := make([]<-chan int, numProcessors)
	for i := 0; i < numProcessors; i++ {
		processors[i] = genCubes(done, inputChan, i+1)
	}
	start := time.Now()
	// aggregator
	for range agrregator(done, processors...) {
		// say we want to recevive only for 200ms
		if time.Since(start).Milliseconds() >= 2000 {
			done <- struct{}{} // signal, so that all the senders stop emiting data and receivers stop expecting data
		}
	}
	fmt.Printf("Elapsed time: %0.5fs\n", time.Since(start).Seconds())

}

func genNums(done chan struct{}, nums ...int) <-chan int {
	out := make(chan int) // unbuffered channel
	go func() {
		defer close(out)
		for _, num := range nums {
			select {
			case out <- num:
			case <-done:
				fmt.Println("cancelling emiting data...") // do some cleanup here
				return
			}
		}
	}()
	return out
}
func genCubes(done chan struct{}, inputChan <-chan int, id int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range inputChan {
			select {
			case out <- num * num * num:
				time.Sleep(fakeProcessTime)
			case <-done:
				return
			}
		}
	}()
	return out
}
func agrregator(done chan struct{}, chs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	output := func(ch <-chan int) {
		for num := range ch {
			select {
			case out <- num:
			case <-done:
				return
			}
		}
	}
	wg.Add(len(chs))
	for _, ch := range chs {
		// start a seperate go routine to handle output from each processor
		go func(ch <-chan int) {
			defer wg.Done()
			output(ch)
		}(ch)
	}
	// routine for close out chan after all the chs are exhausted
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
func fillNums(nums []int, N int) {
	for i := 0; i < N; i++ {
		nums[i] = i + 1
	}
}

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

	inputChan := genNums(nums...) // input stream
	/*
		// dividing the work to two processors or Fan-Out
		cubeChan1 := genCubes(inputChan, 1)
		cubeChan2 := genCubes(inputChan, 2)
		cubeChan3 := genCubes(inputChan, 3)
		// agrregator
		for result := range agrregator(cubeChan1, cubeChan2, cubeChan3) {
			fmt.Println(result)
		}
	*/
	// spawn numProcessors
	processors := make([]<-chan int, numProcessors)
	for i := 0; i < numProcessors; i++ {
		processors[i] = genCubes(inputChan, i+1)
	}
	start := time.Now()
	// aggregator
	for _ = range agrregator(processors...) {
		// fmt.Printf("output:%d\n", result)
	}
	fmt.Printf("Elapsed time: %0.5fs\n", time.Since(start).Seconds())

	/*
	 There is one problem in this approach,
	 in real pipelines, its not always necessay that receivers would receive all the inputs
	 from inbound channels, what if receiver dont want to recevive any further inputs, because of
	 some error in some previous inputs or pipelines. We need to signal the upper pipelines or senders to stop sending
	 This is called explicit cancellation
	*/
}

func genNums(nums ...int) <-chan int {
	out := make(chan int) // unbuffered channel
	go func() {
		for _, num := range nums {
			out <- num // send num
		}
		close(out)
	}()
	return out
}
func genCubes(inputChan <-chan int, id int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range inputChan {
			time.Sleep(fakeProcessTime) // simulate some heavy lifting here
			// fmt.Printf("genCubesId:%d, num:%d\n", id, num)
			out <- num * num * num
		}
		close(out)
	}()
	return out
}
func agrregator(chs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	output := func(ch <-chan int) {
		for num := range ch {
			out <- num
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

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type Vector []int

func (v *Vector) addAll(u *Vector, start int, end int) {
	fmt.Printf("Start:%d, End:%d\n", start, end)
	for i := start; i < end; i++ {
		time.Sleep(3 * time.Second) // fake processing time
		(*v)[i] += (*u)[i]
	}
}

const N = 16

var wg sync.WaitGroup

func main() {
	numCpus := runtime.NumCPU()
	fmt.Printf("Num cpus: %d\n", numCpus)

	v1 := retDummyVector(N)
	// fmt.Printf("v1: %v\n", v1)
	start := time.Now()
	for i := 0; i < numCpus; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			v1.addAll(&v1, i*N/numCpus, (i+1)*N/numCpus) // distributing equal works
		}(i)
	}
	numGoroutines := runtime.NumGoroutine()
	fmt.Printf("Num goroutines: %d\n", numGoroutines)
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Sum DONE\n")
	fmt.Printf("Took: %s\n", &elapsed)
}
func retDummyVector(n int) Vector {
	var vec Vector
	fmt.Printf("generating vectors...\n")
	for i := 0; i < n; i++ {
		vec = append(vec, i)
	}
	fmt.Printf("vector generation done...\n")
	return vec
}

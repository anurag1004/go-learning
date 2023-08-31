package main

import (
	"fmt"
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

func main() {

	v1 := retDummyVector(N)
	start := time.Now()
	v1.addAll(&v1, 0, N)
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

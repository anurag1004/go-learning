package main

import "fmt"

func main() {
	incChan := incrementer(3)
	sumChan := pullAndCalcSum(incChan)
	for sum := range sumChan {
		fmt.Println(sum)
	}
}

func incrementer(N int) chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= N; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func pullAndCalcSum(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int = 0
		for num := range c {
			sum += num
		}
		out <- sum
		close(out)
	}()
	return out
}

package main

import (
	"fmt"
	"strconv"
)

func doSomething(c int, cbk func(int) int, args ...int) {
	fmt.Println("doSomething: start")
	fmt.Println(c)
	fmt.Println(args)
	ret := cbk(c)
	fmt.Println("doSomething: ret=" + strconv.Itoa(ret))
}
func SomeRunner(limit int) int {
	fmt.Println("SomeRunner: ")
	i := 0
	for i < limit {
		fmt.Printf("%v,", i)
		i++
	}
	fmt.Println()
	return i
}

func SomeRetFunc(num int) func(int) int {
	fmt.Println("SomeRetFunc: ")
	return func(count int) int {
		return count * 2 * num
	}
}

// closure
func incrementer() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
func main() {
	fmt.Println("main:")
	args := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	doSomething(10, SomeRunner, args...)

	fmt.Println("main:")
	fn := SomeRetFunc(2)
	fmt.Println(fn(10))

	fmt.Println("main:")
	inc := incrementer()
	i := 0
	for i < 4 {
		i++
		fmt.Println(inc())
	}
}

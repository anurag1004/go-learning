package main

import "fmt"

func foo() {
	defer func() {
		fmt.Println("Defer fooo at start")
	}()
	fmt.Println("foo starts")
	x := []int{0, 1}[0]
	y := 1 / x
	fmt.Println(y)
	defer func() {
		fmt.Println("Defer fooo at end") // this will not excute
	}()
}
func main() {
	foo()
}

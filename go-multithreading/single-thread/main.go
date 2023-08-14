package main

import "fmt"

func main() {
	doSomething()
	doSomethingElse()
}
func doSomething() {
	for i := 0; i < 20; i++ {
		fmt.Printf("doSomething: %d\n", i)
	}
}
func doSomethingElse() {
	for i := 0; i < 20; i++ {
		fmt.Printf("doSomethingElse: %d\n", i)
	}
}

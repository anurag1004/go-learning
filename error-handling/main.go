package main

import (
	"fmt"
	"time"
)

func foo() {
	fmt.Println("foo starts")
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from: %v\n", r)
			fmt.Println("foo ends")
		}
	}()
	bar(0)
	fmt.Println("foo ends")
}
func bar(i int) {
	fmt.Printf("bar %v\n", i)
	if i > 4 {
		fmt.Printf("Panicking at bar(%v)\n", i)
		panic(fmt.Sprintf("bar(%v)", i))
	}
	defer func() {
		fmt.Printf("Defer bar(%v)\n", i)
	}()
	bar(i + 1)
	fmt.Println("bar ends")
}
func main() {
	fmt.Println("Main starts")
	foo()
	time.Sleep(3000) // to demonstrate that we recorvered from panic and our program didnt crashed
	foo()
	fmt.Println("Main ends")
	/*
		Output:
		Main starts
		foo starts
		bar 0
		bar 1
		bar 2
		bar 3
		bar 4
		bar 5
		Panicking at bar(5)
		Defer bar(4)
		Defer bar(3)
		Defer bar(2)
		Defer bar(1)
		Defer bar(0)
		Recovered from: bar(5)
		foo ends
		Main ends
	*/
}

package main

import "fmt"

func c() (i int) {
	// func with a named return value
	// at last 1 will be assigned to i
	return 1

}
func d() (i int) {
	// here we r assigning 21 to i and returning it to named value i
	i = 21
	return
}

func e() (i int) {
	defer func() { i++ }()
	return 1
}
func foo(i int) {
	defer func() {
		fmt.Printf("Defer at foo(%v)\n", i)
	}()
	i++
}
func foo2(i int) {
	defer func(i int) {
		fmt.Printf("Defer at foo(%v)\n", i)
	}(i)
	i++
}
func main() {
	// DEFER function
	/*
		1.) A deferred function’s arguments are evaluated when the defer statement is evaluated.
		2.) Deferred function calls are executed in Last In First Out order after the surrounding function returns.
		3.) Deferred functions may read and assign to the returning function’s named return value
	*/
	fmt.Println(c())
	fmt.Println(d())
	fmt.Println(e())
	foo(10)  // 11
	foo2(10) // 10
}

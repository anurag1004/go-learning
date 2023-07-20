package main

import (
	"fmt"
)

type cell [2]int

func printAnything(i interface{}) {
	fmt.Println(i)
}
func addSomethingtoSlice(myslice *[]int) {
	// slice are always passed by value, copies are made for len, cap and underlying array pointer
	// so any modification inside func will be reflected to the caller.. but when u add an element to slice
	// a new slice is created and this change will not be reflected to the caller
	// because the slice in this func was previously pointing to same underlying array as the caller
	// but when u append an element then this pointer is now pointing to new memory address
	*myslice = append(*myslice, 13)
}
func updateSomthingtoSlice(myslice []int) {
	myslice[0] = 34
}
func main() {
	buff := make([]int, 10)
	str := "hello"
	for _, s := range str {
		fmt.Printf("%c\n", s)
	}
	for i := 0; i < 10; i++ {
		buff[i] = i
	}
	fmt.Println(buff[:4])
	// val := cell{0, 1}
	// fmt.Println(val)
	// printAnything(val)
	// printAnything(12)
	// printAnything(12)
	// printAnything("hello")
	// printAnything('l')
	myslice := []int{10, 11, 12}
	addSomethingtoSlice(&myslice)
	fmt.Println(myslice)
	updateSomthingtoSlice(myslice)
	fmt.Println(myslice)

	b := make([]int, len(myslice)-2)
	copy(b, myslice) // b's len is 2 so,, only two elements from myslice will be copies
	fmt.Println(b)

	// https://go.dev/blog/slices-intro

}

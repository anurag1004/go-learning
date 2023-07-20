package main

import "fmt"

func main() {
	intSlice := []int{1, 2, 3}
	fmt.Println([]float32(intSlice)) // this is a compile time error
	/*
		In go you can not convert a slice of some type to another type. It is not allowed in go
		Although string([]byte{97,67}) this is possible because go language provide this facility for conversion from
		slice of byte to string.
		So can we convert []T1 to []T2 slice if underlying types are same.
		The ans is simply no.
		Consider the below example
	*/
	type T1 int32
	type T2 int32

	t1 := T1(10)
	t2 := T2(20)
	fmt.Println(T1(t2)) // OK
	fmt.Println(T2(t1)) // OK
	var t2Slice []T2
	fmt.Println([]T1(t2Slice)) // compile time error
	/*
		you can think it in two ways..
		1) As mention above we cannot convert one slice type to another
		2) See this structure of slice
		type slice struct {
			array unsafe.Pointer
			len   int
			cap   int
		}
		During convesion the pointer in t2Slice is pointing to address of type t2
		and you are trying to convert to t1, which is simply not possible

		NOTE: This is all true for interface{} or any type. Although u can store anything in any type
		But []interface{}([]int{1,2,3}) is not possible for the same reason as stated above
	*/

}

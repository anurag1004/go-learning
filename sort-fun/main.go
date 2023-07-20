package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name string
	age  int
}
type ByAge []Person

func (by ByAge) Len() int {
	return len(by)
}
func (by ByAge) Less(i, j int) bool {
	return by[i].age < by[j].age
}
func (by ByAge) Swap(i, j int) {
	by[i], by[j] = by[j], by[i]
}
func main() {
	arr := []int{4, 5, 10, 33, 5, 10}
	fmt.Println(arr)
	sort.Ints(arr)
	fmt.Println(arr)

	// Sort is a pack for sorting types in goLang
	/*
		to sort ints.. u can use sort.Ints([]int), similarly for float and other premitive types defined in go
		To sort custom types, there are two methods
		-> sort.Sort(Interface)
		Lets see how Interface is written
		type Interface interface{
			Len() int
			Less(i, j int) bool
			Swap(i,j int)
		}
		Notice ByAge is a type of []Person
		and ByAge also implments all methods of Interface, hence ByAge is also a type of Interface
		This statisfies our sort function
	*/
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	peopleCopy := make([]Person, len(people))
	copy(peopleCopy, people)

	sort.Sort(ByAge(people))
	fmt.Println(people)
	fmt.Println("------------------")
	/*
		This method is lengthy right?
		Lets see another method
		sort provides one more method - Slice(interface{}, func(i,j int) bool)
		the second parameter is our Less method..
	*/
	fmt.Println(peopleCopy)
	sort.Slice(peopleCopy, func(i, j int) bool {
		return peopleCopy[i].age < peopleCopy[j].age
	})
	fmt.Println(peopleCopy)
}

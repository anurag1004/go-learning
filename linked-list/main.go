package main

import (
	"fmt"
)

type node struct {
	next  *node
	value int
}
type list struct {
	head *node
}

func (l list) printList() {
	for n := l.head; n != nil; n = n.next {
		fmt.Printf("%v, ", n.value)
	}
	fmt.Println()
}
func (lst *list) push(val int) {

	if lst.head == nil {
		lst.head = &node{
			next:  nil,
			value: val,
		}
		return
	}
	tempHead := lst.head
	for ; tempHead.next != nil; tempHead = tempHead.next {
	}
	tempHead.next = &node{
		next:  nil,
		value: val,
	}
}
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myLst := list{}
	for _, val := range arr {
		myLst.push(val)
	}
	myLst.printList()
}

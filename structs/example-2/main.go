package main

import "fmt"

type Name struct {
	fname string
	lname string
}

func (n *Name) updateFname(newFname string) {
	n.fname = newFname
}
func (n *Name) updateLname(newLname string) {
	n.fname = newLname
}
func main() {
	name := &Name{
		fname: "Anurag",
		lname: "Verma",
	}
	name.updateFname("Anurag2")
	fmt.Println(name)
}

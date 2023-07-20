package main

type Person struct {
	fname   string
	lname   string
	address Address
}

func (p Person) getFullName() string {
	return p.fname + " " + p.lname
}

func (p *Person) updateFname(newFname string) {
	p.fname = newFname
}
func (p *Person) updateLname(newLname string) {
	p.lname = newLname
}

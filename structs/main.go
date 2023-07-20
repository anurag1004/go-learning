package main

import "fmt"

func main() {
	// if some fields are not provided.. fields are set to nil
	var p Person = Person{
		fname: "Anurag",
		lname: "Verma",
		address: Address{
			stName:  "17th E Cross",
			houseNo: "76",
		},
	}
	fmt.Printf("%+v\n", p)
	ptr := &p
	// fmt.Println((*ptr).getFullName())
	fmt.Println(ptr.getFullName()) // // Go automatically derefrences the pointer
	fmt.Println(ptr.address.getCompleteAddress())
	ptr.updateFname("Anurag_NEW")
	fmt.Println(ptr.getFullName())

	num := 45
	numptr := &num
	*numptr++
	fmt.Println(*numptr, num)

	/*
		Go's is a pass by value lang..
		example: in the receiver func of Person, when we call object's method from the main
		a new copy of that object is created and passed to that funtion.
		So basically we cant really change the fields of the original object.
		To manupilate the value of original object..use references or pointers

		var ptr *Person := &p
		ptr.update()

		p gets copied to receiver func,
		receiver func is defined as *person

		Also, u can basically use the address of the object to call the update method
		(&p).update() -👉 this will work.. but dont use it :) its very wierd looking way to call methods.
		Whenever u use p.method_name(...) and if receiver funtion is a pointer type. Go automatically send the address value of that object.

		Also note - Whenever u pass a reference type object to a method, a new copy in memeory is created in the method block
		basically a new object with SAME values is created i.e address value which the original object was pointing to

		Slices, map, pointers etc are all are reference type
	*/
	fmt.Println((&p).getFullName())
}

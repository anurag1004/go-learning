package main

import (
	"fmt"
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age  uint8
}

func foo(s ...int) {
	fmt.Println(s)
}
func main() {
	p := &Person{
		Name: "Anurag",
		Age:  23,
	}
	myTmpl, err := template.New("test").Parse("Hi {{.Name}}, Your age is {{.Age}}")
	if err != nil {
		panic(err)
	}
	err = myTmpl.Execute(os.Stdout, p)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	list := []int{1, 2, 3, 4}
	foo(list...)
}

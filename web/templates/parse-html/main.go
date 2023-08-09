package main

import (
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age  uint8
}

func main() {
	p := &Person{
		Name: "Anurag",
		Age:  23,
	}
	tmpl, err := template.ParseFiles("test.gohtml")
	if err != nil {
		panic(err)
	}
	file, _ := os.Create("parsed_test.html")
	err = tmpl.Execute(file, p)
	if err != nil {
		panic(err)
	}
}

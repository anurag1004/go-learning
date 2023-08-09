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
	tmpl, _ := template.ParseFiles("test2.gohtml") // returns Template pointer
	// we can add more templates if we want later
	_ = tmpl.ExecuteTemplate(os.Stdout, "test2.gohtml", p)

	tmpl, _ = template.ParseFiles("test3.gohtml")
	_ = tmpl.ExecuteTemplate(os.Stdout, "test3.gohtml", p)
}

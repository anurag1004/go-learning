package main

import (
	"os"
	"text/template"
)

var tmpl *template.Template

type Person struct {
	Name string
	Age  uint8
}

func init() {
	tmpl = template.Must(template.ParseFiles("test.gohtml"))
}
func main() {
	p := &Person{
		Name: "Anurag",
		Age:  23,
	}
	file, _ := os.Create("parsed_test.html")
	err := tmpl.Execute(file, p)
	if err != nil {
		panic(err)
	}
}

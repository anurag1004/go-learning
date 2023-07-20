package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string
	Age  byte
	pin  int // this will be not be included in encoding and decoding json, only upercases will be there
}

func main() {
	person := Person{
		Name: "Anurag",
		Age:  23,
		pin:  4545,
	}
	bs, er := json.Marshal(person)
	if er != nil {
		fmt.Println("Error while encoding Person object")
		os.Exit(-1)
	}
	f, _ := os.Create("people.json")
	bw := bufio.NewWriter(f)
	for _, b := range bs {
		bw.WriteByte(b)
	}
	bw.Flush()

	// using bytes package
	jsonString := bytes.NewBuffer(bs).String()
	fmt.Println(jsonString)

	// using buffered IO
	br := bufio.NewScanner(bytes.NewBuffer(bs))
	for br.Scan() {
		fmt.Println(br.Text())
	}

	var p2 Person
	err := json.Unmarshal(bs, &p2) // bytes to specified type
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Printf("%+v\n", p2)

	// creating object of unkown type
	jsonString2 := `
		{
			"id":12339,
			"date":"23/04/2022",
			"orderStatus":true
		}
	`
	var obj interface{}
	// since every type in Go implements this empty interface, we can cast any unknown json string to Go object
	if err := json.Unmarshal(bytes.NewBufferString(jsonString2).Bytes(), &obj); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println(obj) // it will a map[string]interface{}
	if m, ok := obj.(map[string]interface{}); !ok {
		fmt.Println("map[string]interface{} type not found!")
		os.Exit(-1)
	} else {
		fmt.Printf("{\n")
		for k, v := range m {
			fmt.Printf("\t%v:%v,\n", k, v)
		}
		fmt.Printf("}\n")
	}
}

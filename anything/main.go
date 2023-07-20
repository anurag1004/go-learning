package main

import (
	"encoding/json"
	"fmt"
	"mymodule/anything/horror"
	"os"
)

type Person struct {
	Name string
	Age  byte
	pin  int
	// Here pin will not be parsed by the json encoder
}

func main() {
	horror.GetScreem()

	mp := make(map[string]int)
	mp["anurag"] = 23
	mp["sane"] = 21
	fmt.Printf("%v\n", mp)
	val, ok := mp["anurag"]
	if ok {
		fmt.Printf("found %v\n", val)
	}
	delete(mp, "sane")
	if _, ok := mp["sane"]; !ok {
		fmt.Println("deleted sane")
	}
	// suppose this some json you fetched from an API
	jsonString := `{"Name":"Anurag","Address":"Bangalore","Age":23}`

	// convert the json string to a byte slice
	bs := []byte(jsonString)

	var jsonObject interface{}

	if err := json.Unmarshal(bs, &jsonObject); err != nil {
		panic(err)
	} else {
		if m, ok := jsonObject.(map[string]interface{}); !ok {
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
}

package main

import "fmt"

func main() {
	myMap := make(map[string]string) // equivalent to map[string]string{}
	myMap["id"] = "123"
	myMap["name"] = "anurag"
	fmt.Println(myMap)
	fmt.Println(myMap["age"] == "") // if doesnt exists, zero value of the value type is returned
	_, exists := myMap["age"]       // the second arg tells whether the given key exists in map or not
	fmt.Println(exists)
}

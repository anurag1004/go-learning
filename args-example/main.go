package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("No args provided!")
		os.Exit(-1)
	}
	fileName := os.Args[1]
	fmt.Printf("FileName: %v\n", fileName)
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Read file error:%v\n", err)
		os.Exit(-1)
	}
	os.Stdout.Write(data)
}

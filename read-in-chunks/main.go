package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "file.txt"
	buffSize := 100
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Unable to open file!")
		os.Exit(-1)
	}
	buff := make([]byte, buffSize)
	count := 0
	for {
		byteRead, err := file.Read(buff)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		count++
		if _, err := os.Stat("chunks"); os.IsNotExist(err) {
			err := os.Mkdir("chunks", 0755)
			if err != nil {
				fmt.Println("Unable to create chunks directory!")
				os.Exit(-1)
			}
		}
		fileName := fmt.Sprintf("chunks/chunk-%v.chunk", count)
		chunkFile, err := os.Create(fileName)
		if err != nil {
			fmt.Printf("Error creating chunk file: %v\n", fileName)
			os.Exit(-1)

		}
		numOfBytes, err := chunkFile.Write(buff[:byteRead])
		if err != nil {
			fmt.Printf("Error writing data to chunk file: %v\n", fileName)
			os.Exit(-1)

		}
		fmt.Printf("Bytes Read: %v\n", byteRead)
		fmt.Printf("Bytes wrote to %v, Size: %v\n", fileName, numOfBytes)
		// fmt.Printf("Bytestring to string: %v\n", string(buff[:byteRead]))
	}
}

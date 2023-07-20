package main

import (
	"log"
	"os"
)

func main() {
	log.Println("hi") // 2023/06/17 20:37:11 hi
	// log.Fatal("huhahaha 😈😈😈😈😈") // calls os.Exit(1)
	f, _ := os.Create("test0.log")
	logger := log.New(f, "FROM LOG_EXAMPLE: ", log.Llongfile)
	logger.Println("Hi this is a test log!")
	logger.SetOutput(os.Stdout)
	logger.Fatal("SOME FATAL ERROR")
}

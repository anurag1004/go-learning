package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	// This line of code will read all the data at once and print it..
	// cons: time consuming and high memeory usage
	// temp, _ := os.ReadFile("ebook.txt")
	// fmt.Println(string(temp))

	// Usign buffered IO
	f, err := os.Open("ebook.txt")
	if err != nil {
		log.Fatal("error reading ebook.txt")
	}
	// default size if 4Kb
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	freqMap := make(map[string]int)
	for scanner.Scan() {
		freqMap[scanner.Text()]++
	}
	newF, _ := os.Create("freqcount.txt")
	bw := bufio.NewWriter(newF)
	for key, val := range freqMap {
		bw.WriteString(fmt.Sprintf("%v:%v\n", key, val))
	}
	bw.Flush()
	fmt.Println("Freq file created!")
}

package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const dataFilePath = "data.txt"

func main() {
	initialize(dataFilePath)
	arr := loadFromFile(dataFilePath)
	fmt.Println(arr)
	mergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
func mergeSort(arr []int, i int, j int) {
	if i < j {
		mid := i + (j-i)/2
		mergeSort(arr, i, mid)
		mergeSort(arr, mid+1, j)
		merge(arr, i, mid, j)
	}
}
func merge(arr []int, p int, q int, r int) {
	n1 := q - p + 1
	n2 := r - q

	l1 := make([]int, n1)
	l2 := make([]int, n2)

	for i := 0; i < n1; i++ {
		l1[i] = arr[p+i]
	}
	for i := 0; i < n2; i++ {
		l2[i] = arr[q+1+i]
	}
	var i, j, k int
	k = p
	i = 0
	j = 0
	for i < n1 && j < n2 {
		if l1[i] < l2[j] {
			arr[k] = l1[i]
			i++
		} else {
			arr[k] = l2[j]
			j++
		}
		k++
	}
	for i < n1 {
		arr[k] = l1[i]
		k++
		i++
	}
	for j < n2 {
		arr[k] = l2[j]
		k++
		j++
	}
}
func initialize(filepath string) {
	if len(os.Args) != 2 {
		fmt.Println("Must provide a argument!")
		os.Exit(-1)
	}
	size, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid argument:", os.Args[1])
		os.Exit(-1)
	}
	openOrCreateFile(size, filepath)
}
func openOrCreateFile(size int, dataFilePath string) *os.File {
	var file *os.File
	_, err := os.Stat(dataFilePath)
	if os.IsNotExist(err) {
		// file does not exists.. create one
		file = createFile(dataFilePath)
	} else {
		file, err = os.OpenFile(dataFilePath, os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println("Error opening the existing data file!")
			os.Exit(-1)
		}
		// empty the contents
		file.Truncate(0)
	}
	var sb strings.Builder
	for i := 0; i < size; i++ {
		sb.WriteString(fmt.Sprintf("%v,", getRandomNumber(0, 10000)))
	}
	_, err = io.WriteString(file, sb.String())
	if err != nil {
		fmt.Println("Error writing to file:", err)
		os.Exit(-1)
	}
	fmt.Println("File written successfully")
	return file
}
func createFile(filepath string) *os.File {
	// create a file
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Printf("Error in creating data file: %v Exiting now..\n", filepath)
		os.Exit(-1)
	}
	return file
}
func getRandomNumber(min int, max int) int {
	return int(rand.Float64()*(float64(max-min)) + float64(min))
}
func loadFromFile(filepath string) []int {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("MAIN: Error in reading the data file!")
		os.Exit(1)
	}
	var arr []int
	var buff_size int = 100
	buff := make([]byte, buff_size)
	for {
		byteRead, err := file.Read(buff)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
				os.Exit(1)
			}
			break
		}
		str := string(buff[:byteRead])
		strArr := strings.Split(str, ",")
		for i := 0; i < len(strArr)-1; i++ {
			// n-1 because last there is a comma at last with no following digit
			numStr := strArr[i]
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Printf("%v is not a string!\n", numStr)
			}
			arr = append(arr, num)
		}
	}
	return arr
}

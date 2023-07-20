package main

import (
	"fmt"

	"example.com/mypackage"
	gomyutil2 "github.com/anurag0608/myutil2"
)

func main() {
	// str := "hello"
	fmt.Println(gomyutil2.CapsAllCharArray("helwfefwlo"))
	mypackage.PrintArray([]int{1, 2, 3, 4})
	mypackage.PrintRandomString()
}

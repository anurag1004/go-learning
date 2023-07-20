package mypackage

import (
	"fmt"

	"example.com/mypackage/something"
)

func PrintArray(arr []int) {
	fmt.Println(arr)
}
func PrintRandomString() {
	fmt.Println(something.RandStr())
}
